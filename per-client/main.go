package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hassanjawwad12/per-client/message"
	"github.com/hassanjawwad12/per-client/rate"
)

func endpointHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	resMessage := message.Message{
		Status: "Successful",
		Body:   "Hello, World!",
	}

	err := json.NewEncoder(writer).Encode(&resMessage)
	if err != nil {
		return
	}
}

func main() {

	http.Handle("/ping", rate.PerClientRateLimiter(http.HandlerFunc(endpointHandler)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("There was an error", err)
	}
}
