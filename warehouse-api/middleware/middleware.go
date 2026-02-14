package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"warehouse/config/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Unauthorized(g *gin.Context) {
	g.AbortWithStatusJSON(http.StatusUnauthorized, response.ErrorResponse{
		Success:   false,
		Message:   "Access Denied",
		ErrorCode: "UNAUTHORIZED",
	})
}

func JWTMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(g *gin.Context) {
		authHeader := g.GetHeader("Authorization")
		header := strings.Split(authHeader, " ")

		if len(header) != 2 || header[0] != "Bearer" {
			log.Println("Invalid Authorization header")
			Unauthorized(g)
			return
		}

		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(header[1], claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method")
			}

			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			log.Printf("Token invalid: %v", err.Error())
			Unauthorized(g)
			return
		}

		exp, err := claims.GetExpirationTime()
		if err != nil || time.Now().After(exp.Time) {
			log.Println("Token expired")
			Unauthorized(g)
			return
		}

		userIDFloat, _ := claims["user_id"].(float64)
		userID := int(userIDFloat)
		role, _ := claims["role"].(string)

		g.Set("user_id", userID)
		g.Set("role", role)

		g.Next()
	}
}

func ACLMiddleware(rolesMap map[string]bool) gin.HandlerFunc {
	return func(g *gin.Context) {
		roleVal, exists := g.Get("role")
		if !exists {
			Unauthorized(g)
			return
		}

		role, _ := roleVal.(string)

		if !rolesMap[role] {
			Unauthorized(g)
			return
		}

		g.Next()
	}
}
