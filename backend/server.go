package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pete/go-web/graph"
	"github.com/pete/go-web/middleware"
	"github.com/pete/go-web/service"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize user service
	userService, err := service.NewUserService()
	if err != nil {
		log.Fatalf("Failed to initialize user service: %v", err)
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// CORS middleware
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-User-Id, X-User-Email")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	// Auth middleware
	authMiddleware := middleware.AuthMiddleware(userService)

	// Serve static files from public directory in production
	// Fall back to index.html for client-side routing
	fileServer := http.FileServer(http.Dir("./public"))

	http.Handle("/query", authMiddleware(corsMiddleware(srv)))
	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))

	// Serve static files and SPA fallback
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// If in development mode (no public dir), show playground
		if _, err := os.Stat("./public"); os.IsNotExist(err) {
			playground.Handler("GraphQL playground", "/query").ServeHTTP(w, r)
			return
		}

		// Check if the requested file exists
		path := "./public" + r.URL.Path
		if _, err := os.Stat(path); os.IsNotExist(err) {
			// File doesn't exist, serve index.html for SPA routing
			http.ServeFile(w, r, "./public/index.html")
			return
		}

		// File exists, serve it
		fileServer.ServeHTTP(w, r)
	})

	log.Printf("Server starting on http://localhost:%s", port)
	log.Printf("GraphQL playground: http://localhost:%s/playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
