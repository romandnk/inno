package main

import (
	"context"
	"fmt"
	"inno/attestation1/config"
	"inno/attestation1/internal/entity"
	"inno/attestation1/internal/token"
	"inno/attestation1/internal/worker"
	"inno/attestation1/pkg/storage/inmem"
	"log"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const initialCacheSize uint64 = 100

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	whileList := token.NewWhiteList(cfg.Token)

	cache := inmem.NewCache[string, worker.Data](initialCacheSize)

	// for tests
	input := make(chan entity.Message, 10)

	dir, _ := os.Getwd()
	// random file ids in range [1, 3]
	min, max := 1, 3
	users := make([]entity.Message, 0, 10)
	for i := 0; i < 8; i++ {
		fileId, _ := url.JoinPath(dir, strconv.Itoa(rand.Intn(max-min+1)+min))
		users = append(users, entity.Message{
			Token:  "test1",
			FileID: fileId,
			Data:   "from " + strconv.Itoa(i),
		})
		fmt.Printf("user: %d data: %s\n", i, fileId)
	}
	go func() {
		for _, user := range users {
			input <- user
		}
		//close(input)
	}()

	myWorker := worker.NewWorker(cfg.Worker, cache)
	go myWorker.Run(ctx)

	for {
		select {
		case msg, ok := <-input:
			if !ok {
				return
			}
			whileList.ValidateToken(msg, func(msg entity.Message) {
				cache.Set(ctx, msg.FileID, worker.Data{
					FileID: msg.FileID,
					Data:   msg.Data,
				})
			})
		case <-ctx.Done():
			time.Sleep(cfg.ShutdownTimeout)
			return
		}
	}
}
