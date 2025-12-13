# Multi-stage build for production deployment
# Stage 1: Build frontend
FROM node:22-bookworm-slim AS frontend-builder

WORKDIR /app/frontend

# Copy frontend package files
COPY frontend/package*.json ./

# Install dependencies
RUN npm ci

# Copy frontend source
COPY frontend/ ./

# Build frontend for production
RUN npm run build

# Stage 2: Build backend
FROM golang:1.25-bookworm AS backend-builder

WORKDIR /app

# Copy go mod files
COPY backend/go.mod backend/go.sum ./

# Download dependencies
RUN go mod download

# Copy backend source
COPY backend/ ./

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

# Stage 3: Production image
FROM debian:bookworm-slim

WORKDIR /app

# Install CA certificates for HTTPS requests
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the backend binary
COPY --from=backend-builder /app/server .

# Copy the built frontend files
COPY --from=frontend-builder /app/frontend/dist ./public

# Copy data files if needed
COPY data/ ./data/

# Expose port
EXPOSE 8080

# Set environment variables for production
ENV PORT=8080
ENV DATA_DIR=/app/data

# Run the server
CMD ["./server"]
