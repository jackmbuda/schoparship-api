# Stage 1: The builder
# We use a specific Go version and Alpine Linux for a small base.
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application, creating a static binary.
# CGO_ENABLED=0 is important for creating a static binary that works in a minimal container.
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server .

# Stage 2: The final, minimal image
FROM alpine:latest

WORKDIR /app

# Copy the compiled server binary from the builder stage
COPY --from=builder /app/server .

# Copy the templates and static assets so the server can find them
COPY templates ./templates/
COPY static ./static/

# The application will store its database in a /data directory.
# This directory will be our mount point for the persistent volume.
# We don't need to create it here; Fly.io will handle it.

# Expose port 8080 to the Fly.io proxy
EXPOSE 8080

# The command to run when the container starts
CMD ["/app/server"]