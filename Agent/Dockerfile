
# Use a specific Golang version
FROM golang:1.22.0-alpine as builder

# Set working directory
WORKDIR /app

# Copy your source code
COPY ../Agent .

# Install dependencies
RUN go mod download

# Build the Go binary (replace "main.go" with your actual entry point)
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o agent .

# Use a smaller image for production
FROM alpine:latest

# Set working directory
WORKDIR /app

ENV orchURL="${orchURL}"
ENV MAX_GOROUTINE="${MAX_GOROUTINE}"

# Copy the binary
COPY --from=builder /app/agent /app/agent


# Start the chat API
CMD ["/app/agent"]
