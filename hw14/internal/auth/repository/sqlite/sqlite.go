package sqlite

import (
	"database/sql"
	"github.com/bogatyr285/auth-go/internal/auth/repository/sqlite/token"
	"github.com/bogatyr285/auth-go/internal/auth/repository/sqlite/user"
)

type SQLiteStorage struct {
	db    *sql.DB
	User  *user.SQLLiteUserStorage
	Token *token.SQLLiteTokenStorage
}

func New(dbPath string) (SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return SQLiteStorage{}, err
	}

	userStorage, err := user.New(db)
	if err != nil {
		return SQLiteStorage{}, err
	}

	tokenStorage, err := token.New(db)
	if err != nil {
		return SQLiteStorage{}, err
	}

	return SQLiteStorage{
		db:    db,
		User:  userStorage,
		Token: tokenStorage,
	}, nil
}

func (s *SQLiteStorage) Close() error {
	return s.db.Close()
}
