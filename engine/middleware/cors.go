package middleware

import "net/http"

// CORS middleware
func CORS(allowedOrigins []string, allowedMethods, allowedHeaders []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			// Check if the origin is allowed
			if isOriginAllowed(origin, allowedOrigins) {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			// Set other CORS headers
			w.Header().Set("Access-Control-Allow-Methods", joinStrings(allowedMethods))
			w.Header().Set("Access-Control-Allow-Headers", joinStrings(allowedHeaders))
			w.Header().Set("Access-Control-Allow-Credentials", "true")

			// Handle preflight (OPTIONS) requests
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin || allowedOrigin == "*" {
			return true
		}
	}
	return false
}

func joinStrings(strings []string) string {
	result := ""
	for i, s := range strings {
		result += s
		if i < len(strings)-1 {
			result += ", "
		}
	}
	return result
}
