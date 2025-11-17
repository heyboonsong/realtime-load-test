    # Stage 1: Build the Go application
    FROM golang:1.23-alpine AS builder
    WORKDIR /app
    COPY go.mod go.sum ./
    RUN go mod download
    COPY . .
    RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

    # Stage 2: Create the final image
    FROM alpine:latest
    WORKDIR /app
    COPY --from=builder /app/main .
    EXPOSE 9000
    CMD ["./main"]