package middleware

import (
	"net/http"

	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

var (
	allowedOrigins = []string{"*"}
	allowedMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders = []string{"Content-Type", "Authorization"}
)

func Defaults() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		chimiddleware.Logger,
		chimiddleware.Recoverer,
		chimiddleware.RequestID,
		chimiddleware.AllowContentType("application/json"),
		CORS(allowedOrigins, allowedMethods, allowedHeaders),
	}
}
