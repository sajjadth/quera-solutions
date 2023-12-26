package main

import (
	"fmt"
)

func displayReadableResult(operations []int) string {
	var numberTwo, numberFive int

	for _, o := range operations {
		if o == 2 {
			numberTwo++
		} else {
			numberFive++
		}
	}

	return fmt.Sprintf("5 %v %v", numberTwo+numberFive, -numberTwo)
}

func main() {
	var n int
	var operations []int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var o int
		fmt.Scanln(&o)
		operations = append(operations, o)
	}

	fmt.Println(displayReadableResult(operations))
}
