package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func findSmallestLargerSeries(codes []string) string {
	var output []string
	for _, code := range codes {
		str := []byte(code)

		i := len(str) - 2
		for i >= 0 && str[i] >= str[i+1] {
			i--
		}

		if i == -1 {
			output = append(output, "no answer")
			continue
		}

		j := len(str) - 1
		for str[j] <= str[i] {
			j--
		}

		str[i], str[j] = str[j], str[i]

		sort.Slice(str[i+1:], func(x, y int) bool {
			return str[i+1:][x] < str[i+1:][y]
		})

		output = append(output, string(str))
	}

	return strings.Join(output, "\n")
}

func main() {
	var n int
	var codes []string

	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for i := 0; i < n; i++ {
		scanner.Scan()
		code := scanner.Text()

		codes = append(codes, code)
	}

	fmt.Println(findSmallestLargerSeries(codes))
}
