# Build stage
FROM golang:1.22.3-alpine as builder

# Enable CGO and set environment variables for Alpine
ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Install build tools and SQLite dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the working directory
WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and templates
COPY . .

# Build the Go application
RUN go build -o main cmd/server/main.go

# Run stage
FROM alpine:latest

# Install SQLite runtime dependencies
RUN apk add --no-cache sqlite-libs

WORKDIR /app

# Copy the compiled binary and necessary assets
COPY --from=builder /app/main /app/main
COPY internal/database/migrations /app/internal/database/migrations
COPY templates /app/templates
COPY static /app/static



# Expose the application port
EXPOSE 8082

# Run the application
CMD ["/app/main"]
