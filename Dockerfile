# Build stage
FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./main.go

# Final stage
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/tasks ./tasks

RUN chmod +x /app/main

EXPOSE 8000

CMD ["/app/main"]
