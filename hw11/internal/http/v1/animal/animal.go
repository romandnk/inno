package animal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"net/http"
	"zoo/internal/repository"
)

type Animal struct {
	repo repository.Animal
}

func New(repo repository.Animal) *Animal {
	return &Animal{
		repo: repo,
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

	animal, err := a.repo.GetAnimal(r.Context(), name)
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

	resp := response{
		Animal: animal.Name,
		Amount: animal.Amount,
	}
	err = json.NewEncoder(w).Encode(&resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
