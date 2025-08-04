package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

// mounting to server and requesting for health check -> chi.Mux implements http.Handler
func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // timing out application after taking time to write the server
		ReadTimeout:  time.Second * 10, // timing out when reading from the server *should be less time from writing
		IdleTimeout:  time.Minute,      // timing out when doing nothing or waiting from the server
	}

	log.Printf("server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}
