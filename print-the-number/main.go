package main

import (
	"fmt"
	"strconv"
)

func printNumber(a string) string {
	var output string
	for j, v := range a {
		b, _ := strconv.Atoi(string(v))
		res := fmt.Sprintf("%v: ", b)
		for i := 0; i < b; i++ {
			res = fmt.Sprintf("%v%v", res, b)
		}
		if j == len(a)-1 {
			output += fmt.Sprintf("%v", res)
		} else {
			output += fmt.Sprintf("%v\n", res)

		}
	}
	return output
}

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(printNumber(a))
}
