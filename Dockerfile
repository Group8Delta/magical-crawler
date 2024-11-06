# Stage 1: Build the application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to download dependencies first
COPY go.mod go.sum /app
# RUN go mod download

# Copy the entire project source code to the working directory
COPY . .

# Build the Go application into a binary named "app"
RUN go build -o app ./cmd/magical-crawler

# Stage 2: Create a lightweight image to run the application
FROM alpine:3.18

# Set the working directory for the runtime container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose the application port (adjust as needed)
EXPOSE 8080

# Run the application binary
ENTRYPOINT ["./app"]