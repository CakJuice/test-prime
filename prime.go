/*
 * author Cak Juice
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	var startTime int64 = currentMilliSecond()

	var primes []int
	var limit int = 25000

	for n := 2; n <= limit; n++ {
		var isPrime bool = true

		for div := 2; div < n; div++ {
			if n%div == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, n)
		}
	}

	var endTime int64 = currentMilliSecond()

	fmt.Println(primes)
	fmt.Println("Start Time:", startTime)
	fmt.Println("End Time:", endTime)
	fmt.Println("Selisih:", endTime-startTime, "ms")
}

func currentMilliSecond() int64 {
	now := time.Now()
	return now.UnixNano() / 1000000
}
