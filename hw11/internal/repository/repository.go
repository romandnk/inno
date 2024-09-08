package repository

import (
	"context"
	"zoo/internal/entity"
	animalpostgres "zoo/internal/repository/postgres/animal"
	"zoo/pkg/storage/postgres"
)

type Animal interface {
	GetAnimal(ctx context.Context, animal string) (entity.Animal, error)
}

type Repository struct {
	Animal Animal
}

func New(db *postgres.Postgres) *Repository {
	return &Repository{
		Animal: animalpostgres.New(db),
	}
}
