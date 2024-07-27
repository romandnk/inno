package main

import "sync"

type WorkerPool struct {
	amount int
	jobs   chan string
}

type Result struct {
	Url   string
	Data  []byte
	Error error
}

func NewWorkerPool(workerNum uint64, jobs chan string) WorkerPool {
	return WorkerPool{
		amount: int(workerNum),
		jobs:   jobs,
	}
}

func (wp *WorkerPool) Run(result chan Result, fn func(url string) Result) {
	wg := sync.WaitGroup{}
	wg.Add(wp.amount)
	for range wp.amount {
		go func() {
			defer wg.Done()
			for job := range wp.jobs {
				result <- fn(job)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(result)
	}()
}
