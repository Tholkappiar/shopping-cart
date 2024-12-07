# Stage 1: Build
FROM golang:1.23.4 AS builder

WORKDIR /app

# Copy Go module files and download dependencies
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copy the source code
COPY backend/ ./

# Build the Go binary
RUN go build -o shopping-cart main.go

# Stage 2: Run
FROM ubuntu:jammy

WORKDIR /app

# Install required packages
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Copy the binary from the builder stage
COPY --from=builder /app/shopping-cart .

# Copy the .env file into the container
COPY backend/.env .

# Command to run the application
CMD ["./shopping-cart"]

