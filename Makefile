.PHONY: dev frontend backend install clean help

# Run both frontend and backend concurrently
dev:
	@echo "Starting frontend and backend..."
	@trap 'kill 0' SIGINT; \
	cd frontend && npm run dev & \
	cd backend && air & \
	wait

# Run frontend dev server
frontend:
	@echo "Starting Vue.js frontend..."
	@cd frontend && npm run dev

# Run backend with live reload
backend:
	@echo "Starting Go backend with Air..."
	@cd backend && air

# Install dependencies for both projects
install:
	@echo "Installing frontend dependencies..."
	@cd frontend && npm install
	@echo "Installing backend dependencies..."
	@cd backend && go mod tidy

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

# Show help
help:
	@echo "Available commands:"
	@echo "  make dev       - Run both frontend and backend concurrently"
	@echo "  make frontend  - Run only the frontend dev server"
	@echo "  make backend   - Run only the backend with live reload"
	@echo "  make install   - Install dependencies for both projects"
	@echo "  make build     - Build both frontend and backend"
	@echo "  make clean     - Clean build artifacts"
