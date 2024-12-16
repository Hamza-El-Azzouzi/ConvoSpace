# Build stage
FROM golang:1.22.3-alpine as builder

# Enable CGO for SQLite
ENV CGO_ENABLED=1

# Install build tools and SQLite dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

# Copy Go modules and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and templates
COPY . .

# Build the Go application
RUN go build -o main cmd/server/main.go

# Run stage
# Use a specific version for reproducibility
FROM alpine:3.18 

# Install SQLite runtime dependencies
RUN apk add --no-cache sqlite-libs

WORKDIR /app

# Copy the compiled binary and necessary assets
COPY --from=builder /app/main /app/main
COPY internal/database/migrations /app/internal/database/migrations
COPY templates /app/templates
COPY static /app/static

# Expose the application port (optional)
EXPOSE 8080

# Run the application
CMD ["/app/main"]
