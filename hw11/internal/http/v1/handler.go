package v1

import (
	"net/http"
	"time"
	"zoo/internal/http/v1/animal"
	ratelimiter "zoo/internal/http/v1/middleware/rate_limiter"
	"zoo/internal/repository"
)

type Handler struct {
	mux               *http.ServeMux
	requestNumPerUser int
	rateLimitWindow   time.Duration
	repo              *repository.Repository
}

func NewHandler(requestNumPerUser int, rateLimitWindow time.Duration, repo *repository.Repository) Handler {
	return Handler{
		requestNumPerUser: requestNumPerUser,
		rateLimitWindow:   rateLimitWindow,
		repo:              repo,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	h.mux = http.NewServeMux()

	rateLimiter := ratelimiter.NewIPRateLimiter(h.requestNumPerUser, h.rateLimitWindow)

	h.mux.Handle("/animals/{animal}", rateLimiter.RateLimiter(animal.New(h.repo.Animal)))
	return h.mux
}
