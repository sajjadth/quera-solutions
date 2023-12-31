package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func calculateHIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	for i := 1; i <= len(citations); i++ {
		if i > citations[i-1] {
			return i - 1
		}
	}

	return len(citations)
}

func main() {
	var n int
	var citation []int

	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		scanner.Scan()
		c, _ := strconv.Atoi(scanner.Text())

		citation = append(citation, c)
	}

	fmt.Println(calculateHIndex(citation))
}
