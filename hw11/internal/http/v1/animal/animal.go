package animal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/trace"
	"net/http"
	"zoo/internal/cache"
	"zoo/internal/entity"
	"zoo/internal/repository"
	"zoo/pkg/storage/inmem"
)

type Animal struct {
	repo  repository.Animal
	cache cache.Cache
}

func New(repo repository.Animal, cache cache.Cache) *Animal {
	return &Animal{
		repo:  repo,
		cache: cache,
	}
}

type response struct {
	Animal string `json:"animal"`
	Amount int    `json:"amount"`
}

func (a *Animal) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Создание контекста и span для всего обработчика
	span := trace.SpanFromContext(r.Context())
	defer span.End()

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	name := r.PathValue("animal")

	var animal entity.Animal

	span.AddEvent("try to get animal from cache")
	res, err := a.cache.Get(name)
	if err != nil && !errors.Is(err, inmem.ErrNotFound) {
		span.RecordError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	if res != nil {
		var ok bool
		animal, ok = res.(entity.Animal)
		if !ok {
			err = fmt.Errorf("data from cache is not entity.Animal")
			span.RecordError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
			return
		}
	} else {
		span.AddEvent("get animal from db")
		animal, err = a.repo.GetAnimal(r.Context(), name)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"error": "animal not found"}`))
				return
			}
			span.RecordError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
			return
		}

		go func() {
			_ = a.cache.Set(name, animal)
		}()
	}

	resp := response{
		Animal: animal.Name,
		Amount: animal.Amount,
	}
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		span.RecordError(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
