package ratelimiter

import (
	"net/http"
	"sync"
	"time"
)

type requestInfo struct {
	count        int
	firstAttempt time.Time
}

type IPRateLimiter struct {
	requestNumPerUser int
	rateLimitWindow   time.Duration

	mu      sync.RWMutex
	counter map[string]requestInfo
}

func NewIPRateLimiter(requestNumPerUser int, rateLimitWindow time.Duration) *IPRateLimiter {
	return &IPRateLimiter{
		requestNumPerUser: requestNumPerUser,
		rateLimitWindow:   rateLimitWindow,
		mu:                sync.RWMutex{},
		counter:           make(map[string]requestInfo),
	}
}

func (limiter *IPRateLimiter) RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		now := time.Now().UTC()
		ip := r.RemoteAddr

		info, exists := limiter.counter[ip]
		if !exists {
			limiter.counter[ip] = requestInfo{
				count:        1,
				firstAttempt: now,
			}
		} else {
			if now.Sub(info.firstAttempt) > limiter.rateLimitWindow {
				info.count = 1
				info.firstAttempt = now
			} else {
				if info.count >= limiter.requestNumPerUser {
					w.WriteHeader(http.StatusTooManyRequests)
					return
				}
				info.count++
			}
			limiter.counter[ip] = info
		}

		next.ServeHTTP(w, r)
	})
}
