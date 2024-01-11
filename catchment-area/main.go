package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func adjustStock(i int, side rune, stock, capacity []int) bool {
	n := len(capacity)

	if side == 'L' && i-1 >= 0 && capacity[i-1] < stock[i] {
		stock[i-1] = stock[i]
		if i-1 == 0 || (stock[i-1] > capacity[i-1] && adjustStock(i-1, 'L', stock, capacity)) {
			return true
		}
	} else if side == 'R' && i+1 < n && capacity[i+1] < stock[i] {
		stock[i+1] = stock[i]
		if i+1 == n-1 || (stock[i+1] > capacity[i+1] && adjustStock(i+1, 'R', stock, capacity)) {
			return true
		}
	}

	return false
}

func findFirstOverflowOperation(d, stock, capacity []int, q int) int {
	for i := 0; i < q; i++ {
		j := d[i] - 1
		stock[j] += 1
		if stock[j] > capacity[j] && (adjustStock(j, 'L', stock, capacity) || adjustStock(j, 'R', stock, capacity)) {
			return i + 1
		}
	}

	return 0
}

func main() {
	var n, q int

	fmt.Scan(&n, &q)

	stock := make([]int, n)
	capacity := make([]int, n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i := 0; i < n; i++ {
		scanner.Scan()
		capacity[i], _ = strconv.Atoi(scanner.Text())
	}

	d := make([]int, q)

	for i := 0; i < q; i++ {
		scanner.Scan()
		d[i], _ = strconv.Atoi(scanner.Text())
	}

	res := findFirstOverflowOperation(d, stock, capacity, q)

	if res == 0 {
		fmt.Println("No Answer")
	} else {
		fmt.Println(res)
	}
}
