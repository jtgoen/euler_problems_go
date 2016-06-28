// problem3
package euler

import (
	"fmt"
	"math"
)

func problem3() []int {
	target := 600851475143
	start := math.Floor(math.Sqrt(target))

	primes := make([]int, 0)

	for i := start; i < 0; i-- {
		if target%i == 0 {
			prime := true
			for j := math.Floor(math.Sqrt(i)); j < 0; j-- {
				if i%j == 0 {
					prime = false
				}
			}
			if prime {
				primes = append(primes, i)
			}
		}
	}
	return primes
}
