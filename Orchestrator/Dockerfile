
# Use a specific Golang version
FROM golang:1.22.0-alpine as builder

# Set working directory
WORKDIR /app

# Copy your source code
COPY ../Orchestrator .

# Install dependencies
RUN go mod download

# Build the Go binary (replace "main.go" with your actual entry point)
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o orch ./cmd/main.go

# Use a smaller image for production
FROM alpine:latest

# Set working directory
WORKDIR /app

ENV  DB_NAME="${DB_NAME}"
ENV POSTGRES_USER="${POSTGRES_USER}"
ENV POSTGRES_PASSWORD="${POSTGRES_PASSWORD}"
ENV POSTGRES_HOST="${POSTGRES_HOST}"
ENV POSTGRES_PORT="{POSTGRES_PORT}"

# Copy the binary
COPY --from=builder /app/orch /app/orch

# Expose the port
EXPOSE 8080 

# Start the chat API
CMD ["/app/orch"]
