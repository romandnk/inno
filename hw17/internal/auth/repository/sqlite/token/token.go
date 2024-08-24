package token

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/bogatyr285/auth-go/internal/auth/entity"
	storageerrors "github.com/bogatyr285/auth-go/internal/auth/repository/errors"
	"github.com/bogatyr285/auth-go/internal/auth/usecase"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type SQLLiteTokenStorage struct {
	db *sql.DB
}

func New(db *sql.DB) (*SQLLiteTokenStorage, error) {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS tokens (
		user_id INTEGER not null,
		token TEXT not null,
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
	stmt, err := s.db.PrepareContext(ctx, `INSERT INTO tokens (user_id, token, expiresIn, fingerprint) VALUES(?,?,?,?)`)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(req.UserId, req.Token, req.ExpiresIn, req.Fingerprint); err != nil {
		return err
	}

	return nil
}

func (s *SQLLiteTokenStorage) GetToken(ctx context.Context, token string) (entity.Token, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT user_id, token, expiresIn, fingerprint FROM tokens WHERE token = ?`)
	if err != nil {
		return entity.Token{}, err
	}

	var (
		tk        entity.Token
		expiresIn string
	)
	err = stmt.QueryRow(token).Scan(&tk.UserId, &tk.Token, &expiresIn, &tk.Fingerprint)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.Token{}, storageerrors.ErrTokenNotFound
		}
		return entity.Token{}, err
	}

	tk.ExpiresIn, err = time.Parse("2006-01-02 15:04:05.999999-07:00", expiresIn)
	if err != nil {
		return entity.Token{}, err
	}

	return tk, nil
}

func (s *SQLLiteTokenStorage) DeleteToken(ctx context.Context, token string) error {
	stmt, err := s.db.PrepareContext(ctx, `DELETE FROM tokens WHERE token = ?`)
	if err != nil {
		return err
	}

	if _, err = stmt.Exec(token); err != nil {
		return err
	}

	return nil
}
