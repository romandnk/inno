package authtest

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/bogatyr285/auth-go/cmd/commands"
	"github.com/bogatyr285/auth-go/config"
	"github.com/bogatyr285/auth-go/internal/gateway/http/gen"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
	"time"
)

const configPath = "testdata/config.yaml"

type AuthSuite struct {
	suite.Suite
	cfg    config.Config
	db     *sql.DB
	client *http.Client
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}

func (s *AuthSuite) SetupSuite() {
	cfg, err := config.Parse(configPath)
	if err != nil {
		s.T().Fatalf("failed to parse config: %v", err)
	}
	s.cfg = *cfg

	db, err := sql.Open("sqlite3", cfg.Storage.SQLitePath)
	if err != nil {
		s.T().Fatalf("error parsing config: %s", err.Error())
	}
	s.db = db

	s.client = http.DefaultClient

	cmd := commands.NewServeCmd()
	err = cmd.Flags().Set("config", configPath)
	if err != nil {
		s.T().Fatalf("error adding flag: %s", err.Error())
	}

	go func() {
		if err := cmd.Execute(); err != nil {
			s.T().Fatalf("error starting server: %s", err.Error())
		}
	}()
	// как можно сделать здесь?
	time.Sleep(1 * time.Second)
}

func (s *AuthSuite) TearDownSuite() {
	_, err := s.db.Exec("DROP TABLE IF EXISTS tokens")
	if err != nil {
		s.T().Fatalf("error dropping tokens table: %s", err.Error())
	}

	_, err = s.db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		s.T().Fatalf("error dropping users table: %s", err.Error())
	}

	err = s.db.Close()
	if err != nil {
		s.T().Fatalf("error closing db: %s", err.Error())
	}
}

func (s *AuthSuite) TearDownTest() {
	_, err := s.db.Exec("DELETE FROM tokens")
	if err != nil {
		s.T().Fatalf("error cleaning tokens table: %s", err.Error())
	}
	_, err = s.db.Exec("DELETE FROM users")
	if err != nil {
		s.T().Fatalf("error cleaning users table: %s", err.Error())
	}
}

func (s *AuthSuite) TestPostRegister() {
	existedEmail := "test@test.com"
	_, err := s.db.Exec("INSERT INTO users (email, password) VALUES(?,?)", existedEmail, "test1")
	s.Require().NoError(err)

	requestBody := gen.PostRegisterJSONRequestBody{
		Email:    existedEmail,
		Password: "test2",
	}
	encodedRequestBody, err := json.Marshal(&requestBody)
	s.Require().NoError(err)

	s.T().Log(s.cfg.HTTPServer.Address)
	resp, err := s.client.Post(
		"http://"+s.cfg.HTTPServer.Address+"/register",
		"application/json",
		bytes.NewReader(encodedRequestBody),
	)
	s.Require().NoError(err)

	s.Equal(http.StatusBadRequest, resp.StatusCode)

	var actualResponseBody gen.ErrorResponse
	err = json.NewDecoder(resp.Body).Decode(&actualResponseBody)
	s.Require().NoError(err)

	expectedError := fmt.Sprintf("email already exists: %s", existedEmail)
	s.Equal(expectedError, actualResponseBody.Error)
}
