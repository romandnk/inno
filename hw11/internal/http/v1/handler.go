package v1

import (
	"inno/hw11/internal/http/v1/animal"
	ratelimiter "inno/hw11/internal/http/v1/middleware/rate_limiter"
	"net/http"
	"time"
)

type Handler struct {
	mux               *http.ServeMux
	requestNumPerUser int
	rateLimitWindow   time.Duration
}

func NewHandler(requestNumPerUser int, rateLimitWindow time.Duration) Handler {
	return Handler{
		requestNumPerUser: requestNumPerUser,
		rateLimitWindow:   rateLimitWindow,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	h.mux = http.NewServeMux()

	rateLimiter := ratelimiter.NewIPRateLimiter(h.requestNumPerUser, h.rateLimitWindow)

	h.mux.Handle("/animals/{animal}", rateLimiter.RateLimiter(animal.New()))
	return h.mux
}
