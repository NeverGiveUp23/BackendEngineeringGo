package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("ok")); err != nil {
		return
	}

	app.store.Posts.Create(r.Context())
}
