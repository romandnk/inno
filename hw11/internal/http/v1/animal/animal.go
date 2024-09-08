package animal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
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
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	name := r.PathValue("animal")

	var animal entity.Animal

	res, err := a.cache.Get(name)
	if err != nil && !errors.Is(err, inmem.ErrNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	if res != nil {
		fmt.Println("get from cache")
		var ok bool
		animal, ok = res.(entity.Animal)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
			return
		}
	} else {
		fmt.Println("get from db")
		animal, err = a.repo.GetAnimal(r.Context(), name)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(`{"error": "animal not found"}`))
				return
			}
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
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
