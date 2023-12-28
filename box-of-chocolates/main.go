package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func maxChocolateBoxToAvoidInvitation(boxes, prices []int, k, v, n int) int {
	var maxChocolateBox int

	var chocolateBoxPrice []int

	for i, b := range boxes {
		if b%(k+1) == 1 && b >= k+1 {
			chocolateBoxPrice = append(chocolateBoxPrice, prices[i])
		}
	}

	sort.Slice(chocolateBoxPrice, func(i, j int) bool { return chocolateBoxPrice[i] < chocolateBoxPrice[j] })

	for _, p := range chocolateBoxPrice {
		if p <= v {
			maxChocolateBox++
			v -= p
		} else {
			break
		}
	}

	return maxChocolateBox
}

func main() {
	var k, v, n int

	fmt.Scanln(&k, &v, &n)

	boxes, prices := make([]int, n), make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		boxes[i] = b
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		p, _ := strconv.Atoi(scanner.Text())
		prices[i] = p
	}

	fmt.Println(maxChocolateBoxToAvoidInvitation(boxes, prices, k, v, n))
}
