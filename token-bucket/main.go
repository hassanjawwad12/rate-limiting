package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	message := Message{
		Status: "Successful",
		Body:   "Hello, World!",
	}

	err := json.NewEncoder(writer).Encode(&message)
	if err != nil {
		return
	}
}

func main() {

	http.Handle("/ping", rateLimiter(endpointHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error", err)
	}
}
