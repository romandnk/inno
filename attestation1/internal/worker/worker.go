package worker

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Data struct {
	FileID string
	Data   string
}

type cache interface {
	Set(ctx context.Context, key string, value Data)
	All(ctx context.Context) []Data
	Clear(ctx context.Context)
}

type Worker struct {
	retries int
	num     int
	tick    time.Duration
	cache   cache
}

func New(cfg Config, cache cache) *Worker {
	return &Worker{
		retries: cfg.Retries,
		num:     cfg.Num,
		tick:    cfg.Tick,
		cache:   cache,
	}
}

func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(w.tick)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// get all messages from cache
			messages := w.cache.All(ctx)
			queue := make(chan Data, len(messages))
			go func() {
				for _, m := range messages {
					queue <- m
				}
				close(queue)
			}()
			// Run configured num of workers
			// The problem is that the workers did not record the data during the tick
			// and the next batch was launched and all data for new workers will be cleaned by previous batch.
			// You have to play with the configuration
			wg := &sync.WaitGroup{}
			wg.Add(w.num)
			for workerNum := range w.num {
				go w.writeDataFromQueue(wg, queue, workerNum)
			}
			wg.Wait()
			w.cache.Clear(ctx)
		}
	}
}

func (w *Worker) write(file string, data string) (err error) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			if err != nil {
				err = fmt.Errorf("%w: %w", err, closeErr)
			} else {
				err = closeErr
			}
		}
	}()
	// delete all \n and \t in the beginning and end (I demand it, I am a business 👹)
	data = strings.TrimSpace(data)
	_, err = f.WriteString(data + "\n")
	if err != nil {
		return err
	}
	return nil
}

func (w *Worker) writeDataFromQueue(wg *sync.WaitGroup, queue chan Data, worker int) {
	defer wg.Done()
	for msg := range queue {
		err := w.write(msg.FileID, msg.Data)
		if err != nil {
			log.Printf("(worker %d) failed to write message: %v", worker, err)
			w.doRetry(msg, worker)
		}
	}
}

func (w *Worker) doRetry(data Data, worker int) {
	for retryNum := range w.retries {
		err := w.write(data.FileID, data.Data)
		if err == nil {
			break
		}
		log.Printf("(worker %d) (retry %d) failed to write message: %v", worker, retryNum, err)
	}
}
