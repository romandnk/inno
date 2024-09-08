package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Builder squirrel.StatementBuilderType
	pgxConf *pgxpool.Config
	pool    *pgxpool.Pool
}

func New(cfg Config) (*Postgres, error) {
	pg := Postgres{
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.SSLMode,
	)

	pgxConf, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return &pg, err
	}

	pg.pgxConf = pgxConf

	return &pg, nil
}

func (p *Postgres) GetDB() *pgxpool.Pool {
	return p.pool
}

func (p *Postgres) Connect(ctx context.Context) error {
	db, err := pgxpool.NewWithConfig(ctx, p.pgxConf)
	if err != nil {
		return fmt.Errorf("error creating new pgx pool: %w", err)
	}

	err = db.Ping(ctx)
	if err != nil {
		return fmt.Errorf("error connecting pgx pool: %w", err)
	}

	p.pool = db

	return nil
}

func (p *Postgres) Close() {
	p.pool.Close()
}
