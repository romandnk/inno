package main

import (
	"fmt"
	"math"
	"math/big"
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

	prime, composites := ArrSeparationByPrimeAndCompositeNumbersWithWorkerPool(arr)
	fmt.Println("Primes:", prime)
	fmt.Println("Composites:", composites)
}

func ArrSeparationByPrimeAndCompositeNumbersWithWorkerPool(arr []int) ([]int, []int) {
	numbers := make(chan int, len(arr))

	primesCap, compositesCap := calcChanCapacities(len(arr))

	primes := make(chan int, primesCap)
	composites := make(chan int, compositesCap)

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

	// so we make allocation on most once
	primeNumbers := make([]int, 0, primesCap)
	compositeNumbers := make([]int, 0, compositesCap)

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

func ArrSeparationByPrimeAndCompositeNumbersWithoutWorkerPool(arr []int) ([]int, []int) {
	numbers := make(chan int, len(arr))

	primesCap, compositesCap := calcChanCapacities(len(arr))

	primes := make(chan int, primesCap)
	composites := make(chan int, compositesCap)

	go func() {
		defer close(numbers)
		for i := range arr {
			numbers <- arr[i]
		}
	}()

	go func() {
		for num := range numbers {
			if big.NewInt(int64(num)).ProbablyPrime(0) {
				primes <- num
			} else {
				composites <- num
			}
		}
		close(primes)
		close(composites)
	}()

	// so we make allocation on most once
	primeNumbers := make([]int, 0, primesCap)
	compositeNumbers := make([]int, 0, compositesCap)

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

// calcChanCapacities calculates the amount of prime and composite numbers respectively
// formula is taken from https://en.wikipedia.org/wiki/Prime_number_theorem
func calcChanCapacities(amount int) (int, int) {
	primeAmount := int(math.Ceil(float64(amount) / math.Log(float64(amount))))
	compositeAmount := amount - primeAmount
	return primeAmount, compositeAmount
}
