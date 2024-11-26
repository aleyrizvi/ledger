package engine

import "net/http"

type Route struct {
	Handler    func(w http.ResponseWriter, r *http.Request)
	Method     string
	Path       string
	Middleware []func(http.Handler) http.Handler
}

type RouterConfig struct {
	Router  http.Handler
	Pattern string
}
