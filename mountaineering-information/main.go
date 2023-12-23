package main

import "fmt"

func minDistanceFromAminsImaginaryCastle(distances []int) int {
	var total, countTwo, countOne int

	for _, d := range distances {
		if d == 1 {
			countOne++
		} else if d == 2 {
			countTwo++
		}
		total += d
	}
	if countOne < countTwo && countOne%2 == 0 && countTwo%2 == 1 {
		return 2
	}

	return total % 2
}

func main() {
	var n int
	var distances []int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var d int
		fmt.Scan(&d)
		distances = append(distances, d)
	}

	fmt.Println(minDistanceFromAminsImaginaryCastle(distances))
}
