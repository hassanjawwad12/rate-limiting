package rate

import (
	"encoding/json"
	"net/http"

	"github.com/hassanjawwad12/token-bucket/message"
	"golang.org/x/time/rate"
)

func RateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	limiter := rate.NewLimiter(2, 4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// not allowing requests if the limiter is full
		if !limiter.Allow() {

			resMessage := message.Message{
				Status: "Error",
				Body:   "API is at its capacity, try again later",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&resMessage)
			return
		} else {
			// call the endpoint handler in the main file
			next(w, r)
		}
	})
}
