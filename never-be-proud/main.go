package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func chromaticNumberOfGraph(data [][2]int) int {
	var output int
	res := make([]int, len(data)*2+2)

	for _, d := range data {
		res[d[0]]++
		res[d[1]]++

		if output < res[d[0]] {
			output = res[d[0]]
		}
		if output < res[d[1]] {
			output = res[d[1]]
		}
	}

	return output
}

func main() {
	var n, v, u int
	fmt.Scanln(&n)

	data := make([][2]int, n-1)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < n-1; i++ {
		scanner.Scan()
		v, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		u, _ = strconv.Atoi(scanner.Text())

		data[i] = [2]int{v, u}
	}

	fmt.Println(chromaticNumberOfGraph(data))
}
