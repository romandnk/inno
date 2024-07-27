package test

import (
	"context"
	"github.com/stretchr/testify/suite"
	"inno/attestation1/config"
	"inno/attestation1/internal/entity"
	"inno/attestation1/internal/token"
	"inno/attestation1/internal/worker"
	"inno/attestation1/pkg/storage/inmem"
	"inno/attestation1/test/testdata"
	"os"
	"strconv"
	"testing"
	"time"
)

type Suite struct {
	suite.Suite
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestSuccess() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const (
		whiteToken string        = "test"
		userNum    int           = 10
		tickTime   time.Duration = 250 * time.Millisecond
	)

	cfg := config.Config{
		Worker: worker.Config{
			Num:     5,
			Tick:    tickTime,
			Retries: 2,
		},
		Token: token.Config{
			WhiteTokens: []string{whiteToken},
		},
		ShutdownTimeout: time.Millisecond * 125,
	}

	whileList := token.NewWhiteList(cfg.Token)

	cache := inmem.NewCache[string, worker.Data](uint64(userNum))

	files, err := testdata.CreateTempFiles(suite.T().TempDir(), userNum)
	suite.Require().NoError(err)

	input := make(chan entity.Message, userNum)
	messages := make([]entity.Message, userNum)
	for i := 0; i < userNum; i++ {
		msg := testdata.GenerateUserRequest(strconv.Itoa(i), whiteToken, files[i])
		messages[i] = msg
		input <- msg
	}
	close(input)

	for msg := range input {
		whileList.ValidateToken(msg, func(msg entity.Message) {
			cache.Set(ctx, msg.FileID, worker.Data{
				FileID: msg.FileID,
				Data:   msg.Data,
			})
		})
	}

	myWorker := worker.NewWorker(cfg.Worker, cache)
	go myWorker.Run(ctx)

	time.Sleep(tickTime * 2)

	for i, file := range files {
		data, err := os.ReadFile(file.Name())
		suite.Require().NoError(err)

		suite.Require().Equal(messages[i].Data+"\n", string(data))
	}

	suite.Require().Len(cache.All(ctx), 0)
}

func (suite *Suite) TestInvalidToken() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const (
		whiteToken string        = "test"
		userNum    int           = 10
		tickTime   time.Duration = 250 * time.Millisecond
	)

	cfg := config.Config{
		Worker: worker.Config{
			Num:     5,
			Tick:    tickTime,
			Retries: 2,
		},
		Token: token.Config{
			WhiteTokens: []string{whiteToken},
		},
		ShutdownTimeout: time.Millisecond * 125,
	}

	whileList := token.NewWhiteList(cfg.Token)

	cache := inmem.NewCache[string, worker.Data](uint64(userNum))

	files, err := testdata.CreateTempFiles(suite.T().TempDir(), userNum)
	suite.Require().NoError(err)

	input := make(chan entity.Message, userNum)
	messages := make([]entity.Message, userNum)
	for i := 0; i < userNum; i++ {
		var msg entity.Message
		if i%2 == 0 {
			msg = testdata.GenerateUserRequest(strconv.Itoa(i), whiteToken, files[i])
		} else {
			msg = testdata.GenerateUserRequest(strconv.Itoa(i), "invalid", files[i])
		}

		messages[i] = msg
		input <- msg
	}
	close(input)

	for msg := range input {
		whileList.ValidateToken(msg, func(msg entity.Message) {
			cache.Set(ctx, msg.FileID, worker.Data{
				FileID: msg.FileID,
				Data:   msg.Data,
			})
		})
	}

	// valid messages will be from userId 0, 2, 4, 6, 8
	suite.Require().Len(cache.All(ctx), 5)

	myWorker := worker.NewWorker(cfg.Worker, cache)
	go myWorker.Run(ctx)

	time.Sleep(tickTime * 2)

	for i, file := range files {
		data, err := os.ReadFile(file.Name())
		suite.Require().NoError(err)

		if i%2 == 0 {
			suite.Require().Equal(messages[i].Data+"\n", string(data))
		} else {
			suite.Require().Empty(string(data))
		}
	}

	suite.Require().Len(cache.All(ctx), 0)
}
