package main

import "fmt"

func calculateMinUncoverableLength(a, b int) int {
	return a % b
}

func main() {
	var a, b int

	fmt.Scanln(&a, &b)

	fmt.Println(calculateMinUncoverableLength(a, b))
}
