.PHONY: dev dev-with-auth frontend backend install clean build generate docker-build docker-run docker-stop nginx-proxy help

# Run both frontend and backend concurrently
dev:
	@echo "Starting frontend and backend..."
	@trap 'kill 0' SIGINT; \
	cd frontend && npm run dev & \
	export DATA_DIR="$$(pwd)/data" && cd backend && air & \
	wait

# Run frontend and backend with nginx proxy for auth testing
dev-with-auth:
	@echo "Starting frontend, backend, and nginx proxy..."
	@echo "Access the app at http://localhost:3001 (as admin user)"
	@trap 'kill 0' SIGINT; \
	cd frontend && npm run dev & \
	export DATA_DIR="$$(pwd)/data" && cd backend && air & \
	nginx -c "$$(pwd)/nginx.conf" -g 'daemon off;' & \
	wait

# Run nginx proxy (requires frontend and backend to be running)
nginx-proxy:
	@echo "Starting nginx proxy on port 3001..."
	@echo "Make sure frontend and backend are running first!"
	@nginx -c "$$(pwd)/nginx.conf" -g 'daemon off;'

# Run frontend dev server
frontend:
	@echo "Starting Vue.js frontend..."
	@cd frontend && npm run dev

# Run backend with live reload
backend:
	@echo "Starting Go backend with Air..."
	@if [ -f .env ]; then export $$(cat .env | grep -v '^#' | xargs); fi && export DATA_DIR="$$(pwd)/data" && cd backend && air

# Install dependencies for both projects
install:
	@echo "Installing frontend dependencies..."
	@cd frontend && npm install
	@echo "Installing backend dependencies..."
	@cd backend && go mod download
	@cd backend && go mod tidy
	@echo "Installing gqlgen..."
	@go install github.com/99designs/gqlgen@latest
	@echo "Installing air for live reloading..."
	@go install github.com/air-verse/air@latest

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@cd frontend && rm -rf node_modules dist
	@cd backend && rm -rf tmp

# Build both projects
build:
	@echo "Building frontend..."
	@cd frontend && npm run build
	@echo "Building backend..."
	@cd backend && go build -o bin/server .

# Regenerate GraphQL code
generate:
	@echo "Regenerating GraphQL code..."
	@cd backend && gqlgen generate
	@cd backend && go mod tidy

# Build Docker image for production
docker-build:
	@echo "Building production Docker image..."
	@docker build -t go-web-app:latest .
	@echo "Build complete! Run with: make docker-run"

# Run Docker container
docker-run:
	@echo "Starting Docker container in background on port 8080..."
	@docker run -d -p 8080:8080 -v $(PWD)/data:/app/data --name go-web-app go-web-app:latest
	@echo "Container started! Access at http://localhost:8080"
	@echo "View logs: docker logs go-web-app -f"
	@echo "Stop container: make docker-stop"

# Stop and remove Docker container
docker-stop:
	@echo "Stopping Docker container..."
	@docker stop go-web-app 2>/dev/null || true
	@docker rm go-web-app 2>/dev/null || true
	@echo "Container stopped and removed."

# Show help
help:
	@echo "Available commands:"
	@echo "  make dev                 - Run both frontend and backend concurrently"
	@echo "  make dev-with-auth       - Run frontend, backend, and nginx proxy for auth testing"
	@echo "  make nginx-proxy         - Run nginx proxy (requires frontend/backend running)"
	@echo "  make frontend            - Run only the frontend dev server"
	@echo "  make backend             - Run only the backend with live reload"
	@echo "  make install             - Install dependencies for both projects"
	@echo "  make build               - Build both frontend and backend"
	@echo "  make generate            - Regenerate GraphQL code from schema"
	@echo "  make clean               - Clean build artifacts"
	@echo "  make docker-build        - Build production Docker image"
	@echo "  make docker-run          - Run Docker container (detached)"
	@echo "  make docker-stop         - Stop and remove Docker container"
