package token

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bogatyr285/auth-go/internal/auth/usecase"
	_ "github.com/mattn/go-sqlite3"
)

type SQLLiteTokenStorage struct {
	db *sql.DB
}

func New(db *sql.DB) (*SQLLiteTokenStorage, error) {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS tokens (
		user_id INTEGER not null,
		refresh TEXT not null,
		expiresIn timestamptz not null,
		fingerprint TEXT not null
	);
	create index if not exists idx_username ON users(username);
	`)
	if err != nil {
		return nil, fmt.Errorf("error init table 'tokens': %w", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("error create table 'tokens': %w", err)
	}

	return &SQLLiteTokenStorage{db: db}, nil
}

func (s *SQLLiteTokenStorage) SaveToken(ctx context.Context, req usecase.SaveTokenRequest) error {
	stmt, err := s.db.PrepareContext(ctx, `INSERT INTO tokens (user_id, refresh, expiresIn, fingerprint) VALUES(?,?,?,?)`)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(req.UserId, req.Token, req.ExpiresIn, req.Fingerprint); err != nil {
		return err
	}

	return nil
}
