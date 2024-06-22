package main

import (
	"fmt"
	"sync"
)

const numsChannels int = 2

/*3. Реализуйте функцию слияния двух каналов в один.*/

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	for i := 0; i < 3; i++ {
		ch1 <- i
		ch2 <- i * i
	}
	close(ch1)
	close(ch2)

	res := mergeChannels(ch1, ch2)

	for val := range res {
		fmt.Println(val)
	}
}

func mergeChannels(ch1, ch2 chan int) chan int {
	out := make(chan int, len(ch1)+len(ch2))

	wg := sync.WaitGroup{}
	wg.Add(numsChannels)
	go func() {
		defer wg.Done()
		for val := range ch1 {
			out <- val
		}
	}()
	go func() {
		defer wg.Done()
		for val := range ch2 {
			out <- val
		}
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
