package transversal_domain

import (
	"fmt"
	"net/http"
)

// Middleware interface defines the Handle method for chaining
type Middleware interface {
	Handle(http.Handler) http.Handler
}

func HandlerWrapper(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// set headers for view json content type
		w.Header().Set("Content-Type", "application/json")

		handler(w, r)
	}
}

// LoggingMiddleware logs incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Incoming request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Use ServeHTTP for cleaner chaining
	})
}

// AuthenticationMiddleware checks for valid credentials (replace with your logic)
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate checking for a valid token in the request header
		if r.Header.Get("X-Auth-Token") != "secret123" {
			http.Error(w, "Unauthorized access", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// AuthorizationMiddleware checks for specific user roles (replace with your logic)
func AuthorizationMiddleware(next http.Handler, allowedRoles []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate checking for a user role in the request context
		userRole := r.Context().Value("userRole")
		if userRole == nil {
			http.Error(w, "User role not found", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
		// Consider setting user info in context here for later access
	})
}
