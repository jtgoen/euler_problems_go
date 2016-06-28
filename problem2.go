// problem2
package euler

import (
	"fmt"
)

func fibonacci(n int, max int) int {
	return 0
}

func problem2 int {
	evenFibTotal := 0

	prev := 0
	for i := 1; i <= 4000000; {
		temp := i
		i += prev
		prev = temp

		if i%2 == 0 {
			evenFibTotal += i
		}
	}

	return evenFibTotal
}
