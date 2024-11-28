package main

import (
	"fmt"
	"log/slog"

	"github.com/aleyrizvi/ledger/config"
	"github.com/aleyrizvi/ledger/engine"
	"github.com/aleyrizvi/ledger/engine/middleware"
	"github.com/aleyrizvi/ledger/postgres"
	"github.com/aleyrizvi/ledger/user"
)

func main() {
	c := config.New()

	dbrw, dbro := postgres.New(c.DBRW, c.DBRW)
	defer func() {
		dbrw.Close()
		dbro.Close()
	}()

	userRepository := user.NewRepository(dbrw, dbro)
	userService := user.NewService(userRepository)
	fmt.Println(userService)
	userHandler := user.NewHandler(userService)

	e := engine.New(&engine.Config{
		Logger:      slog.Default(),
		Middlewares: middleware.Defaults(),
		Routes: []engine.RouterConfig{
			{
				Pattern: "/users/",
				Router:  userHandler,
			},
		},
		HTTPPort: c.Port,
	})

	e.Run()
}
