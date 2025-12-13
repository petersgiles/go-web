# Dev Container Dockerfile
FROM mcr.microsoft.com/devcontainers/go:1.25-bookworm

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - \
	&& apt-get install -y nodejs

# Install Go tools as vscode user
USER vscode
RUN go install github.com/99designs/gqlgen@latest \
	&& go install github.com/air-verse/air@latest

# Ensure Go cache directory has correct permissions
RUN mkdir -p /go/pkg/mod/cache && chmod -R 755 /go/pkg/mod

# Set working directory
WORKDIR /app

# Expose ports
EXPOSE 3001 5173 8080
