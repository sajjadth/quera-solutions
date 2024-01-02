package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func countInversions(arr [4]int) float64 {
	res := 0.0

	for j := 1; j < 4; j++ {
		for i := j + 1; i < 4; i++ {
			if arr[i] > arr[j] {
				res++
			}
		}
	}

	return res
}

func minPassengersDiscomfort(destinations []int) int {
	minDiscomfort := math.MaxInt

	for _, perm := range [][]int{{0, 1, 2, 3}, {1, 0, 2, 3}, {2, 0, 1, 3}, {3, 0, 1, 2}} {
		arrPerm := [4]int{destinations[perm[0]], destinations[perm[1]], destinations[perm[2]], destinations[perm[3]]}
		minDiscomfort = int(math.Min(float64(minDiscomfort), float64(countInversions(arrPerm))))
	}

	return minDiscomfort
}

func main() {
	var n int
	var output []string

	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for i := 0; i < n; i++ {
		var destinations []int

		for j := 0; j < 4; j++ {
			scanner.Scan()
			destination, _ := strconv.Atoi(scanner.Text())

			destinations = append(destinations, destination)
		}

		output = append(output, strconv.Itoa(minPassengersDiscomfort(destinations)))
	}

	fmt.Println(strings.Join(output, "\n"))
}
