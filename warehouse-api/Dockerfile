FROM golang:1.24.6-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8000

CMD ["./app"]