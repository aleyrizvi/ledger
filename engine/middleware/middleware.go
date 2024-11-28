package middleware

import (
	"net/http"
)

var (
	allowedOrigins     = []string{"*"}
	allowedMethods     = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	allowedHeaders     = []string{"Content-Type", "Authorization"}
	allowedSourceTypes = []string{"game", "server", "payment"}
)

func Defaults() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		CORS(allowedOrigins, allowedMethods, allowedHeaders),
		AllowedSourceTypes(allowedSourceTypes),
	}
}
