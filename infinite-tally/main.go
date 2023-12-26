package main

import (
	"fmt"
	"math"
)

func findMinimumStages(m, n, a, b int) int {
	minBoard, maxEraser := m, a

	if n < m {
		minBoard = n
	}
	if a < b {
		maxEraser = b
	}

	return int(math.Ceil(float64(minBoard) / float64(maxEraser)))
}

func main() {
	var m, n, a, b int

	fmt.Scanln(&n)
	fmt.Scanln(&m)
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	fmt.Println(findMinimumStages(m, n, a, b))
}
