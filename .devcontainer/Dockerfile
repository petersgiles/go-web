# Dev Container Dockerfile
FROM mcr.microsoft.com/devcontainers/go:1.25-bookworm

# Install Node.js
RUN curl -fsSL https://deb.nodesource.com/setup_22.x | bash - \
	&& apt-get install -y nodejs

# Install typst
WORKDIR /tmp
RUN curl -fsSL https://github.com/typst/typst/releases/download/v0.14.2/typst-x86_64-unknown-linux-musl.tar.xz -o typst.tar.xz \
	&& tar -xJf typst.tar.xz \
	&& find . -name typst -type f -exec mv {} /usr/local/bin/typst \; \
	&& chmod +x /usr/local/bin/typst \
	&& rm typst.tar.xz
WORKDIR /app
USER vscode
RUN go install github.com/99designs/gqlgen@latest \
	&& go install github.com/air-verse/air@latest

# Ensure Go cache directory has correct permissions
RUN mkdir -p /go/pkg/mod/cache && chmod -R 755 /go/pkg/mod

# Set working directory
WORKDIR /app

