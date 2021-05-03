// package main

// import (
// 	"encoding/json"
// 	"net/http"
// 	"fmt"

// 	"github.com/go-chi/chi/v5"
// 	"github.com/go-chi/chi/v5/middleware"
// )

// type Sample struct {
// 	Message string `json: "message"`
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	sample := Sample{Message: "hello-chi"}
// 	json.NewEncoder(w).Encode(sample)
	
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
//     fmt.Println(sample)
//     fmt.Println("Endpoint Hit: homePage")
// }

// func handleRequests() {
// 	r := chi.NewRouter()
// 	r.Use(middleware.Logger)
// 	r.Get("/", homePage)
// 	http.ListenAndServe(":8080", r)
// }

// func main() {
// 	handleRequests()
// }