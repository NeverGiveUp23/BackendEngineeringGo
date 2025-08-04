package main

import (
	"BackendEngineeringGo/internal/store"
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("ok")); err != nil {
		return
	}

	err := app.store.Posts.Create(r.Context(), &store.Post{})
	if err != nil {
		return
	}
}
