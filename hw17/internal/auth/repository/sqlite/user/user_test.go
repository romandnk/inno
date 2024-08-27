package user

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bogatyr285/auth-go/internal/auth/entity"
	storageerrors "github.com/bogatyr285/auth-go/internal/auth/repository/errors"
	"github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type UserStorageSuite struct {
	suite.Suite
	ctx         context.Context
	mock        sqlmock.Sqlmock
	userStorage *SQLLiteUserStorage
}

func TestUserStorageSuite(t *testing.T) {
	suite.Run(t, new(UserStorageSuite))
}

// SetupSubTest runs before each subtest in the suite
func (s *UserStorageSuite) SetupSubTest() {
	db, mock, err := sqlmock.New()
	s.Require().NoError(err)

	query := regexp.QuoteMeta(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		email text not null unique,
		password text not null,
		created_at timestamptz DEFAULT CURRENT_TIMESTAMP);
	create index if not exists idx_username ON users(email);
	`)
	mock.ExpectPrepare(query).
		WillBeClosed().
		ExpectExec().
		WillReturnResult(sqlmock.NewResult(0, 1))

	// конечно не бест практис делать такие миграции в коде приложения, но пусть будут)
	userStorage, err := New(db)
	s.Require().NoError(err)

	s.ctx = context.Background()
	s.mock = mock
	s.userStorage = userStorage
}

// TearDownSubTest runs after each subtest in the suite
func (s *UserStorageSuite) TearDownSubTest() {
	err := s.mock.ExpectationsWereMet()
	s.Require().NoError(err)
}

func (s *UserStorageSuite) TestUserStorage_RegisterUser() {
	query := regexp.QuoteMeta(`INSERT INTO users (email, password) VALUES(?,?) RETURNING id`)

	testCases := []struct {
		name         string
		expectedUser entity.UserAccount
		mock         func()
		expectedErr  error
		expectedId   int
	}{
		{
			name: "success",
			expectedUser: entity.UserAccount{
				Email:    "test@test.com",
				Password: "1234",
			},
			mock: func() {
				s.mock.ExpectPrepare(query).
					ExpectQuery().
					WithArgs("test@test.com", "1234").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			expectedErr: nil,
			expectedId:  1,
		},
		{
			name: "nickname already exists",
			expectedUser: entity.UserAccount{
				Email:    "test@test.com",
				Password: "1234",
			},
			mock: func() {
				s.mock.ExpectPrepare(query).
					ExpectQuery().
					WithArgs("test@test.com", "1234").
					WillReturnError(sqlite3.Error{
						Code: sqlite3.ErrConstraint,
					})
			},
			expectedErr: storageerrors.ErrNicknameAlreadyExists,
			expectedId:  0,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			tc.mock()

			id, err := s.userStorage.RegisterUser(s.ctx, tc.expectedUser)
			if tc.expectedErr != nil {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
			}

			s.Equal(tc.expectedId, id)
		})
	}
}
