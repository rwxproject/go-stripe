package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler{
	mux := chi.NewRouter()
	mux.Get("/virtual-terminal", app.VirtualTerminal)
	return mux
}