package main

import (
	"fmt"
	"strconv"
)

func swapFirstAndLastDigits(n int) int {
	digits := strconv.Itoa(n)
	if len(digits) <= 1 {
		return n
	}

	firstDigit, lastDigit := string(digits[0]), string(digits[len(digits)-1])
	reversedDigits := lastDigit + digits[1:len(digits)-1] + firstDigit

	result, _ := strconv.Atoi(reversedDigits)
	return result
}

func compare(a, b int) string {
	reversedA := swapFirstAndLastDigits(a)
	reversedB := swapFirstAndLastDigits(b)

	switch {
	case reversedA == reversedB:
		return fmt.Sprintf("%v = %v", a, b)
	case reversedA < reversedB:
		return fmt.Sprintf("%v < %v", a, b)
	default:
		return fmt.Sprintf("%v < %v", b, a)
	}
}

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(compare(a, b))
}
