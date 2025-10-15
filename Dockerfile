# Build stage
FROM golang:1.24-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git make

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application and migration tool
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrate ./cmd/migrations

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata postgresql-client

WORKDIR /root/

# Copy the binaries from builder
COPY --from=builder /app/server .
COPY --from=builder /app/migrate .

# Copy config example (will be overridden by environment variables)
COPY --from=builder /app/configs/config.yaml.example ./configs/config.yaml

# Copy migrations
COPY --from=builder /app/migrations ./migrations

# Create entrypoint script
RUN echo '#!/bin/sh' > /root/entrypoint.sh && \
    echo 'set -e' >> /root/entrypoint.sh && \
    echo 'echo "Waiting for database..."' >> /root/entrypoint.sh && \
    echo 'until pg_isready -h ${APP_DATABASE_HOST:-localhost} -p ${APP_DATABASE_PORT:-5432} -U ${APP_DATABASE_USER:-postgres}; do' >> /root/entrypoint.sh && \
    echo '  sleep 2' >> /root/entrypoint.sh && \
    echo 'done' >> /root/entrypoint.sh && \
    echo 'echo "Database is ready!"' >> /root/entrypoint.sh && \
    echo 'echo "Running migrations..."' >> /root/entrypoint.sh && \
    echo 'ls -la /root/migrations || echo "Migrations directory not found"' >> /root/entrypoint.sh && \
    echo './migrate -dir=/root/migrations -v up || echo "Migration failed, but continuing..."' >> /root/entrypoint.sh && \
    echo 'echo "Starting server..."' >> /root/entrypoint.sh && \
    echo 'exec ./server' >> /root/entrypoint.sh && \
    chmod +x /root/entrypoint.sh

# Expose port
EXPOSE 8080

# Run the entrypoint script
CMD ["/root/entrypoint.sh"]
