# Authentication Setup

This application uses a federated identity approach where authentication is handled by a proxy (nginx) that injects user headers into requests.

## Architecture

1. **Federated Identity**: Authentication is handled externally by a reverse proxy (nginx)
2. **User Profile**: User profiles are stored in `data/users.json` and loaded by the application
3. **Header-based Auth**: User identity is passed via HTTP headers (`X-User-Id`, `X-User-Email`)
4. **Dev Default**: In development, nginx always injects a default dev user

## Development Setup

Start the development environment:

```bash
make dev
```

This starts (via Docker Compose):
- Frontend on port 5173
- Backend on port 8080
- Nginx proxy on port 3001 (with dev user auth)

**Access the app at http://localhost:3001** (nginx proxies and adds auth headers)

Default logged-in user: **Dev User** (`dev-user@example.com`)

To stop:
```bash
make dev-stop
```

### Testing Different Users

Edit `nginx.conf` and change the `X-User-Id` header to test different users:

Available test users (from `data/users.json`):
- **Dev User** (default): `dev-user` / `dev-user@example.com`
- **Admin User**: `user-001` / `admin@example.com`
- **John Doe**: `user-002` / `john.doe@example.com`
- **Jane Smith**: `user-003` / `jane.smith@example.com`
- **Bob Wilson**: `user-004` / `bob.wilson@example.com`

Then restart the proxy:

```bash
docker-compose restart nginx-proxy
```

### Development Without Docker

To run backend/frontend directly (for faster iteration, no auth):

```bash
make backend  # or make frontend
```

Without nginx, `currentUser` returns `null`.

## GraphQL API

### Query Current User

```graphql
query {
  currentUser {
    id
    email
    name
    roles
    avatar
    department
    permissions
  }
}
```

Returns `null` if no user is authenticated (no headers provided).

## Production Setup

In production, use a reverse proxy or API gateway to:

1. Authenticate users (OAuth, SAML, etc.)
2. Set appropriate headers:
   - `X-User-Id`: User's unique identifier
   - `X-User-Email`: User's email address

The Go backend will:
1. Extract headers via `middleware/auth.go`
2. Load user profile from `data/users.json` (or database in production)
3. Add user to request context
4. Make user available to GraphQL resolvers

## Code Structure

- `backend/middleware/auth.go` - Auth middleware that extracts user from headers
- `backend/service/user_service.go` - User profile service
- `backend/graph/schema.graphqls` - GraphQL schema with User type
- `backend/graph/schema.resolvers.go` - currentUser resolver
- `data/users.json` - User profiles
- `nginx.conf` - Development proxy configuration
