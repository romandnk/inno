package main

import (
	"math/big"
	"sync"
)

type WorkerPool struct {
	Nums       chan int
	Primes     chan int
	Composites chan int
	wg         *sync.WaitGroup
}

func NewWorkerPool(amount int, nums chan int, primes, composites chan int) *WorkerPool {
	wg := &sync.WaitGroup{}
	wg.Add(amount)
	go func() {
		wg.Wait()
		close(primes)
		close(composites)
	}()
	return &WorkerPool{
		Nums:       nums,
		Primes:     primes,
		Composites: composites,
		wg:         wg,
	}
}

func (pool *WorkerPool) Work() {
	go func() {
		defer pool.wg.Done()
		for num := range pool.Nums {
			if big.NewInt(int64(num)).ProbablyPrime(0) {
				pool.Primes <- num
			} else {
				pool.Composites <- num
			}
		}
	}()
}
