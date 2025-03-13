package message

import (
	"time"

	"golang.org/x/time/rate"
)

type Message struct {
	Status string `json:"status"`
	Body   string `json:"body"`
}

type Client struct {
	Limiter  *rate.Limiter
	LastSeen time.Time
}
