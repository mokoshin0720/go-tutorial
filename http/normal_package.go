package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Sample struct {
	Message string `json:"message"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	sample := Sample{Message: "hello"}
	
    res, err := json.Marshal(sample)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(res)
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}