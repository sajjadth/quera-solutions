package main

import "fmt"

func minTotalShoutsForDaraAndSara(n int) int {
	if n == 1 {
		return 2
	} else {
		return 3
	}
}

func main() {
	var n int

	fmt.Scanln(&n)

	fmt.Println(minTotalShoutsForDaraAndSara(n))
}
