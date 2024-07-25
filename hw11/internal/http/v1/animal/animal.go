package animal

import (
	"encoding/json"
	"net/http"
)

type Auth struct {
	animals map[string]int
}

func New() *Auth {
	return &Auth{
		animals: map[string]int{
			"elephants":  5,
			"monkeys":    3,
			"crocodiles": 6,
		},
	}
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	animal := r.PathValue("animal")

	amount, ok := a.animals[animal]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response := struct {
		Animal string `json:"animal"`
		Amount int    `json:"amount"`
	}{
		Animal: animal,
		Amount: amount,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
