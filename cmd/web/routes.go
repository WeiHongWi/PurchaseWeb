package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	MUX := chi.NewRouter()

	return MUX
}
