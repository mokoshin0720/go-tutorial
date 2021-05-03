package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Sample struct {
	Message string `json: "message"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	sample := Sample{Message: "hello"}
	json.NewEncoder(w).Encode(sample)

    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
