package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/didip/tollbooth/v8"
	"github.com/didip/tollbooth/v8/limiter"
	"github.com/hassanjawwad12/toolbooth/message"
)

func endpointHandler(writer http.ResponseWriter, request *http.Request) {

	time.Sleep(100 * time.Millisecond)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	resMessage := message.Message{
		Status: "Successful",
		Body:   "Hello, World!",
	}

	err := json.NewEncoder(writer).Encode(&resMessage)
	if err != nil {
		log.Println("Error encoding response:", err)
		return
	}
}

func main() {

	toolboothLimiter := tollbooth.NewLimiter(1, nil)

	// Explicitly define how to pick the IP address
	toolboothLimiter.SetIPLookup(limiter.IPLookup{
		Name:           "X-Forwarded-For",
		IndexFromRight: 0,
	})

	resMsg := message.Message{
		Status: "success",
		Body:   "Hello World",
	}
	jsonMsg, _ := json.Marshal(resMsg)

	toolboothLimiter.SetMessageContentType("application/json")
	toolboothLimiter.SetMessage(string(jsonMsg))

	http.Handle("/ping", tollbooth.HTTPMiddleware(toolboothLimiter)(http.HandlerFunc(endpointHandler)))
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error starting server: ", err)
		return
	}
}
