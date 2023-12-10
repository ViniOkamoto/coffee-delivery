package application

import (
	"context"
	"net/http"
)

type App struct {
	router http.Handler
}

func NewApp() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

func (a *App) Run(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
