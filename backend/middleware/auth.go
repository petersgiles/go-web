package middleware

import (
	"context"
	"net/http"

	"github.com/pete/go-web/service"
)

// contextKey is a type for context keys to avoid collisions
type contextKey string

const (
	// UserContextKey is the key for storing user in context
	UserContextKey contextKey = "user"
	// UserIDContextKey is the key for storing user ID in context
	UserIDContextKey contextKey = "userID"
)

// AuthMiddleware extracts user information from headers and adds it to the context
func AuthMiddleware(userService *service.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// Extract user ID from header (set by nginx or other auth proxy)
			userID := r.Header.Get("X-User-Id")
			if userID == "" {
				// Also try email-based lookup
				userEmail := r.Header.Get("X-User-Email")
				if userEmail != "" {
					if user, found := userService.GetUserByEmail(userEmail); found {
						userID = user.ID
					}
				}
			}

			// If we have a user ID, fetch the user profile and add to context
			if userID != "" {
				ctx = context.WithValue(ctx, UserIDContextKey, userID)

				if user, found := userService.GetUserByID(userID); found {
					ctx = context.WithValue(ctx, UserContextKey, user)
				}
			}

			// Continue with the updated context
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
