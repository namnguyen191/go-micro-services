package main

import (
	"net/http"
)

func (app *Config) Broker(rw http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "hit broker",
	}

	_ = app.writeJSON(rw, http.StatusOK, payload)
}
