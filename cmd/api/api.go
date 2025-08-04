package main

import (
	"BackendEngineeringGo/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	MaxIdleTime  string
}

// mounting to server and requesting for health check -> chi.Mux implements http.Handler
func (app *application) mount() http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from panic

	r.Use(middleware.Timeout(60 * time.Second))

	// simple api versioning
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
