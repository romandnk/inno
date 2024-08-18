package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	storageerrors "github.com/bogatyr285/auth-go/internal/auth/repository/errors"
	"github.com/mattn/go-sqlite3"

	"github.com/bogatyr285/auth-go/internal/auth/entity"
	_ "github.com/mattn/go-sqlite3"
)

type SQLLiteUserStorage struct {
	db *sql.DB
}

func New(db *sql.DB) (*SQLLiteUserStorage, error) {
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		email text not null unique,
		password text not null,
		created_at timestamptz DEFAULT CURRENT_TIMESTAMP);
	create index if not exists idx_username ON users(email);
	`)
	if err != nil {
		return nil, fmt.Errorf("error init table 'users': %w", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("error create table 'users': %w", err)
	}

	return &SQLLiteUserStorage{db: db}, nil
}

func (s *SQLLiteUserStorage) RegisterUser(ctx context.Context, u entity.UserAccount) (int, error) {
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
				return 0, storageerrors.ErrNicknameAlreadyExists
			}
		}

		return 0, err
	}

	return id, nil
}

func (s *SQLLiteUserStorage) FindUserByEmail(ctx context.Context, email string) (entity.UserAccount, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT id, password FROM users WHERE email = ?`)
	if err != nil {
		return entity.UserAccount{}, err
	}

	var (
		userId     int
		pswdFromDB string
	)

	if err := stmt.QueryRow(email).Scan(&userId, &pswdFromDB); err != nil {
		return entity.UserAccount{}, err
	}

	return entity.UserAccount{
		Id:       userId,
		Email:    email,
		Password: pswdFromDB,
	}, nil
}
