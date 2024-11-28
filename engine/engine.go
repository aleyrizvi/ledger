package engine

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

const (
	readTimeOut  = 3 * time.Second
	writeTimeOut = 3 * time.Second
)

// Config allows configuration for Engine constructor
type Config struct {
	Logger      *slog.Logger
	Middlewares []func(http.Handler) http.Handler
	Routes      []RouterConfig
	HTTPPort    uint
}

type Engine struct {
	server *http.Server
	logger *slog.Logger
}

func New(c *Config) *Engine {
	e := &Engine{}

	if c.Logger != nil {
		e.logger = c.Logger
	} else {
		e.logger = slog.Default()
	}

	e.logger.With("context", "engine.New").Info("New engine Initializing")

	// create router
	r := http.NewServeMux()

	// // mount all the sub routers
	for _, route := range c.Routes {
		handler := route.Router
		for _, mw := range c.Middlewares {
			handler = applyMiddleware(handler, mw)
		}
		r.Handle(route.Pattern, http.StripPrefix(strings.TrimSuffix(route.Pattern, "/"), handler))
	}

	e.server = &http.Server{
		Addr:         fmt.Sprintf(":%d", c.HTTPPort),
		ReadTimeout:  readTimeOut,
		WriteTimeout: writeTimeOut,
		Handler:      r,
	}

	return e
}

func (e *Engine) Run() {
	if err := e.server.ListenAndServe(); err != nil {
		e.logger.With("context", "engine.Run").Error(err.Error())
	}
}
