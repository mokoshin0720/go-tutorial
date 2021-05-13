// controllerをroutingに登録する。server.goから呼び出される
package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func some(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "これで完璧？")
    fmt.Println("Endpoint Hit: homePage")
}

func HandleRequests() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	r.Get("/", some)

	return r
}