package main

import (
	"fmt"
	"strings"
)

func checkMatchingNecklace(necklaces [2]string) string {
	if len(necklaces[0]) != len(necklaces[1]) {
		return "NO"
	}

	concatenated := necklaces[0] + necklaces[0]

	for i := 0; i < len(necklaces[0]); i++ {
		rotated := concatenated[i : i+len(necklaces[0])]
		if rotated == necklaces[1] || reverse(rotated) == necklaces[1] {
			return "YES"
		}
	}

	return "NO"
}

func reverse(s string) string {
	var result strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		result.WriteByte(s[i])
	}
	return result.String()
}

func main() {
	var t int
	var outputs []string

	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		var a, b string
		fmt.Scanln(&a, &b)
		outputs = append(outputs, checkMatchingNecklace([2]string{a, b}))
	}

	fmt.Println(strings.Join(outputs, "\n"))
}
