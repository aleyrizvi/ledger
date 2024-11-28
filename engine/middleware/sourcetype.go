package middleware

import (
	"net/http"
)

func AllowedSourceTypes(allowedSourceTypes []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sourceType := r.Header.Get("Source-Type")
			if sourceType == "" {
				http.Error(w, "Source-Type header is required", http.StatusBadRequest)
				return
			}

			for _, t := range allowedSourceTypes {
				if t == sourceType {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Source-Type header is invalid", http.StatusBadRequest)
		})
	}
}
