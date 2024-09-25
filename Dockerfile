# Use the official Golang image to build the Go app
FROM golang:1.22-alpine AS builder

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create a directory for the app
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the WebSocket Go application
RUN go build -o main .

# Final stage: start with a small base image for the production
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary built in the previous stage
COPY --from=builder /app/main .

# Expose the port your WebSocket server listens on (e.g., 8080)
EXPOSE 8080

# Run the Go app
CMD ["./main"]
