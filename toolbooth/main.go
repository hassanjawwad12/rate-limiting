package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/didip/tollbooth/v8"
	"github.com/didip/tollbooth/v8/limiter"
	"github.com/hassanjawwad12/toolbooth/message"
)

func endpointHandler(writer http.ResponseWriter) {
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

func logRateLimitFailure(w http.ResponseWriter, r *http.Request) {
	// Get client IP address
	clientIP := r.RemoteAddr
	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		clientIP = strings.Split(forwarded, ",")[0]
	}

	log.Printf("Rate limit exceeded for IP: %s", clientIP)

	// Respond with JSON error message
	response := message.Message{
		Status: "error",
		Body:   "Too many requests. Please try again later.",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusTooManyRequests)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Set rate limit: 1 request per second
	toolboothLimiter := tollbooth.NewLimiter(3, nil)

	// Use RemoteAddr for local testing
	toolboothLimiter.SetIPLookup(limiter.IPLookup{Name: "RemoteAddr"})

	// Define custom response message for rate limit errors
	resMsg := message.Message{
		Status: "error",
		Body:   "Too many requests. Please try again later.",
	}
	jsonMsg, _ := json.Marshal(resMsg)
	toolboothLimiter.SetMessageContentType("application/json")
	toolboothLimiter.SetMessage(string(jsonMsg))

	// Middleware function to log rate limit failures
	http.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpError := tollbooth.LimitByRequest(toolboothLimiter, w, r)
		if httpError != nil {
			logRateLimitFailure(w, r)
			return
		}
		endpointHandler(w)
	}))

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Error starting server:", err)
	}
}
