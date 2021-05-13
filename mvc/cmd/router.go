// controllerをroutingに登録する。server.goから呼び出される
package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
	"tutorial/mvc/pkg/controller"
)

func HandleRequests() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	r.Get("/", controller.Index)
	r.Post("/create", controller.Create)

	return r
}