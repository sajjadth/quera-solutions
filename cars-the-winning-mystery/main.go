package main

import "fmt"

func getWinnerCarNumber(loop, length float64, carsInfo [][2]float64) int {
	var winnerTime float64
	var winnerNumber int

	for i, v := range carsInfo {
		currentTime := v[0] + ((length * loop) / v[1])
		currentNumber := i + 1

		if i == 0 || winnerTime > currentTime {
			winnerTime, winnerNumber = currentTime, currentNumber
		}
	}

	return winnerNumber
}

func main() {
	var loop, length float64
	var n int
	var carsInfo [][2]float64

	fmt.Scanln(&loop, &length)
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var t, v float64

		fmt.Scanln(&t, &v)

		carsInfo = append(carsInfo, [2]float64{t, v})
	}

	fmt.Println(getWinnerCarNumber(loop, length, carsInfo))
}
