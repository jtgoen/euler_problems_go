// problem1
package euler

import (
	"fmt"
)

func problem1 int {
	total := 0

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return total
}
