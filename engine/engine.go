package engine

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

const (
	readTimeOut  = 3 * time.Second
	writeTimeOut = 3 * time.Second
)

// Config allows configuration for Engine constructor
type Config struct {
	Logger     *slog.Logger
	Middleware []func(http.Handler) http.Handler
	Route      []RouterConfig
	HTTPPort   uint
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

	// create chi router
	r := chi.NewRouter()

	// apply all global middlewares. These are applied at the root fo the router and will run
	// before any routes
	r.Use(c.Middleware...)

	// mount all the sub routers
	for _, route := range c.Route {
		r.Mount(route.Pattern, route.Router)
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
	e.server.ListenAndServe()
}
