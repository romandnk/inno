package v1

import (
	"net/http"
	"time"
	"zoo/internal/cache"
	"zoo/internal/http/v1/animal"
	ratelimiter "zoo/internal/http/v1/middleware/rate_limiter"
	"zoo/internal/repository"
)

type Handler struct {
	mux               *http.ServeMux
	requestNumPerUser int
	rateLimitWindow   time.Duration
	repo              *repository.Repository
	cache             cache.Cache
}

func NewHandler(requestNumPerUser int, rateLimitWindow time.Duration, repo *repository.Repository, cache cache.Cache) Handler {
	return Handler{
		requestNumPerUser: requestNumPerUser,
		rateLimitWindow:   rateLimitWindow,
		repo:              repo,
		cache:             cache,
	}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	h.mux = http.NewServeMux()

	rateLimiter := ratelimiter.NewIPRateLimiter(h.requestNumPerUser, h.rateLimitWindow)

	h.mux.Handle("/animals/{animal}", rateLimiter.RateLimiter(animal.New(h.repo.Animal, h.cache)))
	return h.mux
}
