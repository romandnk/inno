package animalpostgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"zoo/internal/entity"
	"zoo/pkg/storage/postgres"
)

type Repo struct {
	db *postgres.Postgres
}

func New(db *postgres.Postgres) *Repo {
	return &Repo{db}
}

func (r *Repo) GetAnimal(ctx context.Context, name string) (entity.Animal, error) {
	query, args, err := r.db.Builder.
		Select("name", "amount").
		From("animals").
		Where(squirrel.Eq{"name": name}).
		ToSql()
	if err != nil {
		return entity.Animal{}, fmt.Errorf("error creating query: %w", err)
	}

	var animal entity.Animal
	err = r.db.GetDB().QueryRow(ctx, query, args...).Scan(&animal.Name, &animal.Amount)
	if err != nil {
		return entity.Animal{}, fmt.Errorf("error getting animal: %w", err)
	}

	return animal, nil
}
