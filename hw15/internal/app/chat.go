package app

import (
	"chat/internal/handler"
	"chat/internal/repository/cache"
	"chat/internal/server"
	"chat/internal/service"
	"chat/pkg/authclient"
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
)

func Run() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	var wg sync.WaitGroup

	chatDB, err := cache.ChatCacheInit(ctx, &wg)
	if err != nil {
		log.Fatalf("ERROR failed to initialize chat database: %v", err)
	}
	// initialize service
	service.Init(chatDB)
	authclient.Init("localhost:8000")

	go func() {
		err := server.Run("localhost", "8001", http.HandlerFunc(handler.HandleHTTPReq))
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("ERROR server run ", err)
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(); err != nil {
		log.Fatal("ERROR server shutdown ", err)
	}
	wg.Wait()
	log.Println("INFO chat service was gracefully shutdown")
}
