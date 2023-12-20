package main

import "fmt"

func printHappinessStatus(x1, x2, x3, x4, y1, y2, y3, y4 int) string {
	var a int
	if x1 == x3 || x1 == x4 || y1 == y3 || y1 == y4 {
		a++
	}
	if x2 == x3 || x2 == x4 || y2 == y3 || y2 == y4 {
		a++
	}
	if a%2 == 0 {
		return "unhappy"
	} else {
		return "happy"
	}
}

func main() {
	var x1, x2, x3, x4, y1, y2, y3, y4 int

	fmt.Scanln(&x1, &y1)
	fmt.Scanln(&x2, &y2)
	fmt.Scanln(&x3, &y3)
	fmt.Scanln(&x4, &y4)

	fmt.Println(printHappinessStatus(x1, x2, x3, x4, y1, y2, y3, y4))
}
