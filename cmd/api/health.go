package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	//if _, err := w.Write([]byte(`{"status": "ok"}`)); err != nil {
	//	return
	//}

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env, // development, staging
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "err.Error()")
	}
}
