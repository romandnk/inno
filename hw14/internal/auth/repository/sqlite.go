package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"

	"github.com/bogatyr285/auth-go/internal/auth/entity"
	_ "github.com/mattn/go-sqlite3"
)

var ErrNicknameAlreadyExists = errors.New("nickname already exists")

type SQLLiteStorage struct {
	db *sql.DB
}

func New(dbPath string) (SQLLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return SQLLiteStorage{}, err
	}
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		email text not null unique,
		password text not null,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
	create index if not exists idx_username ON users(email);
	`)
	if err != nil {
		return SQLLiteStorage{}, fmt.Errorf("error init table 'users': %w", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return SQLLiteStorage{}, fmt.Errorf("error create table 'users': %w", err)
	}

	stmt, err = db.Prepare(`
	CREATE TABLE IF NOT EXISTS tokens (
		id INTEGER PRIMARY KEY,
		username text not null unique,
		password text not null,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
	create index if not exists idx_username ON users(username);
	`)
	if err != nil {
		return SQLLiteStorage{}, fmt.Errorf("error init table 'tokens': %w", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return SQLLiteStorage{}, fmt.Errorf("error create table 'tokens': %w", err)
	}

	return SQLLiteStorage{db: db}, nil
}

func (s *SQLLiteStorage) Close() error {
	return s.db.Close()
}

func (s *SQLLiteStorage) RegisterUser(ctx context.Context, u entity.UserAccount) (int, error) {
	stmt, err := s.db.PrepareContext(ctx, `INSERT INTO users (email, password) VALUES(?,?) RETURNING id`)
	if err != nil {
		return 0, err
	}

	var id int
	err = stmt.QueryRow(u.Email, u.Password).Scan(&id)
	if err != nil {
		var sqliteErr sqlite3.Error
		if errors.As(err, &sqliteErr) {
			if errors.Is(sqliteErr.Code, sqlite3.ErrConstraint) {
				return 0, ErrNicknameAlreadyExists
			}
		}

		return 0, err
	}

	return id, nil
}

func (s *SQLLiteStorage) FindUserByEmail(ctx context.Context, email string) (entity.UserAccount, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT password FROM users WHERE email = ?`)
	if err != nil {
		return entity.UserAccount{}, err
	}

	var pswdFromDB string

	if err := stmt.QueryRow(email).Scan(&pswdFromDB); err != nil {
		return entity.UserAccount{}, err
	}

	return entity.UserAccount{
		Email:    email,
		Password: pswdFromDB,
	}, nil
}
