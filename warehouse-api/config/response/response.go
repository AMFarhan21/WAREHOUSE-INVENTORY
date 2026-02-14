package response

type Meta struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}

const (
	InsufficientStock = "INSUFFICIENT_STOCK"
	ItemNotFound      = "ITEM_NOT_FOUND"
	ValidationError   = "VALIDATION_ERROR"
	InternalError     = "INTERNAL_ERROR"
)
