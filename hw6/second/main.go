package main

import (
	"fmt"
	"sync"
	"time"
)

func RunProcessor(wg *sync.WaitGroup, mu *sync.Mutex, prices []*sync.Map) {
	go func() {
		defer wg.Done()

		mu.Lock()
		defer mu.Unlock()
		for _, price := range prices {
			price.Range(func(k, v any) bool {
				p, ok := v.(float64)
				if !ok {
					return false
				}
				price.Store(k, p+1)

				return true
			})
			price.Range(func(k, v any) bool {
				fmt.Print(k, " ", v, " ")

				return true
			})
			fmt.Printf("\n")
		}
	}()
}

func RunWriter() <-chan *sync.Map {
	var prices = make(chan *sync.Map)
	go func() {

		var currentPrice = sync.Map{}
		currentPrice.Store("inst1", 1.1)
		currentPrice.Store("inst2", 2.1)
		currentPrice.Store("inst3", 3.1)
		currentPrice.Store("inst4", 4.1)

		for i := 1; i < 5; i++ {
			currentPrice.Range(func(k, v any) bool {
				price, ok := v.(float64)
				if !ok {
					return false
				}
				currentPrice.Store(k, price+1)
				return true
			})
			temp := CopySyncMap(currentPrice)
			prices <- &temp
			time.Sleep(time.Second)
		}
		close(prices)
	}()
	return prices
}
func main() {
	p := RunWriter()

	mu := sync.Mutex{}
	var prices []*sync.Map

	for price := range p {
		prices = append(prices, price)
	}

	for _, price := range prices {
		price.Range(func(k, v any) bool {
			fmt.Print(k, " ", v, " ")

			return true
		})
		fmt.Printf("\n")
	}
	fmt.Println()

	wg := sync.WaitGroup{}
	wg.Add(3)
	RunProcessor(&wg, &mu, prices)
	RunProcessor(&wg, &mu, prices)
	RunProcessor(&wg, &mu, prices)
	wg.Wait()
}

func CopySyncMap(m sync.Map) sync.Map {
	var cp sync.Map

	m.Range(func(k, v any) bool {
		vm, ok := v.(sync.Map)
		if ok {
			cp.Store(k, CopySyncMap(vm))
		} else {
			cp.Store(k, v)
		}

		return true
	})

	return cp
}
