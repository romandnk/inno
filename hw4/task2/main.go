package main

import (
	"fmt"
	"sync"
)

/*2. Напишите функцию разделения массива чисел на массивы простых и составных чисел.
Для записи в массивы используйте два разных канала и горутины.
Важно, чтобы были использованы владельцы каналов.*/

const workerPoolNum int = 3

func main() {
	arr := make([]int, 1000)
	for i := range 1000 {
		arr[i] = i
	}

	prime, composites := arrSeparationByPrimeAndCompositeNumbers(arr)
	fmt.Println("Primes:", prime)
	fmt.Println("Composites:", composites)
}

func arrSeparationByPrimeAndCompositeNumbers(arr []int) ([]int, []int) {
	numbers := make(chan int, len(arr))

	primes := make(chan int, len(arr)/5)   // amount of prime numbers will always be too small in relation to len(arr)
	composites := make(chan int, len(arr)) // amount of composites numbers will always be less than len(arr)

	pool := NewWorkerPool(workerPoolNum, numbers, primes, composites)

	for i := 0; i < workerPoolNum; i++ {
		pool.Work()
	}

	go func() {
		defer close(numbers)
		for i := range arr {
			numbers <- arr[i]
		}
	}()

	primeNumbers := make([]int, 0, len(arr)/5)
	compositeNumbers := make([]int, 0, len(arr))

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for num := range primes {
			primeNumbers = append(primeNumbers, num)
		}
	}()

	go func() {
		defer wg.Done()
		for num := range composites {
			compositeNumbers = append(compositeNumbers, num)
		}
	}()
	wg.Wait()

	return primeNumbers, compositeNumbers
}
