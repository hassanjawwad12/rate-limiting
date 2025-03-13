package rate

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/hassanjawwad12/per-client/message"
	"golang.org/x/time/rate"
)

func PerClientRateLimiter(next func(w http.ResponseWriter, r *http.Request)) http.Handler {

	var (
		mu      sync.Mutex
		clients = make(map[string]*message.Client)
	)

	// remove clients from the struct after 2 minutes of inactivity
	go func() {
		for {
			time.Sleep(1 * time.Minute)
			mu.Lock()
			for ip, client := range clients {
				if time.Since(client.LastSeen) > 2*time.Minute {
					delete(clients, ip)
				}
			}
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		mu.Lock()
		if _, found := clients[ip]; !found {
			// Client is calling the api for the first time
			clients[ip] = &message.Client{Limiter: rate.NewLimiter(2, 4)}
		}
		clients[ip].LastSeen = time.Now()
		// if limit is reached
		if !clients[ip].Limiter.Allow() {
			mu.Unlock()
			resMessage := message.Message{
				Status: "Error",
				Body:   "API is at its capacity, try again later",
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(&resMessage)
			return
		}
		mu.Unlock()
		next(w, r)
	})
}
