package main

import (
	"fmt"
	"math"
)

func calculateMrFirouzPaymentAmount(cakes []int, k int) int {
	var min int

	if k == 1 {
		for _, v := range cakes {
			if min < v {
				min = v
			}
		}
	} else if k == 2 {
		if len(cakes) > 2 {
			if cakes[0] < cakes[len(cakes)-1] {
				min = cakes[0]
			} else {
				min = cakes[len(cakes)-1]
			}
		} else {
			min = cakes[0]
		}
	} else {
		min = math.MaxInt
		for _, v := range cakes {
			if v < min {
				min = v
			}
		}
	}

	return min
}

func main() {
	var n, k int
	var cakes []int

	fmt.Scanln(&n, &k)

	for i := 0; i < n; i++ {
		var c int
		fmt.Scan(&c)
		cakes = append(cakes, c)
	}

	fmt.Println(calculateMrFirouzPaymentAmount(cakes, k))
}
