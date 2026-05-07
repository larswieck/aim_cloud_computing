# Build stage — compiles the binary inside the Go toolchain image
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o deploy-tracker .

# Runtime stage — copies only the binary into a minimal image
FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/deploy-tracker .
ENTRYPOINT ["./deploy-tracker"]
