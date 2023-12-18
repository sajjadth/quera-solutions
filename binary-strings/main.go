package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func operationRequired(strs []string) int {
	rows, cols := len(strs), len(strs[0])
	result := 0
	colCounts := make([]int, cols)

	for _, s := range strs {
		for j := 0; j < cols-1; j++ {
			if s[j] != s[j+1] {
				colCounts[j]++
			}
		}

		if s[cols-1] == '1' {
			colCounts[cols-1]++
		}
	}

	for _, count := range colCounts {
		result += int(math.Min(float64(count), float64(rows-count)))
	}

	return result
}

func main() {
	var t int
	var results []string

	fmt.Scanln(&t)

	for i := 0; i < t; i++ {
		var n, m int
		var strs []string

		fmt.Scanln(&n, &m)

		for j := 0; j < n; j++ {
			var str string
			fmt.Scanln(&str)
			strs = append(strs, str)
		}

		results = append(results, strconv.Itoa(operationRequired(strs)))
	}

	fmt.Println(strings.Join(results, "\n"))
}
