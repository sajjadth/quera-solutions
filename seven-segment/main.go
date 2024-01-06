package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sevenSegment(v string) int {
	var res int
	power := map[int]int{0: 6, 1: 2, 2: 5, 3: 5, 4: 4, 5: 5, 6: 6, 7: 3, 8: 7, 9: 6}

	num := strings.Split(strings.Split(v, "e")[0], ".")
	zero, _ := strconv.Atoi(strings.Split(v, "e")[1])

	if len(num) == 1 {
		res += power[0] * zero
	} else {
		res += power[0] * (zero - len(num[1]))
	}
	for _, v := range strings.Split(strings.Join(num, ""), "") {
		n, _ := strconv.Atoi(v)
		res += power[n]
	}

	return res
}

func main() {
	var n string
	fmt.Scanln(&n)
	fmt.Println(sevenSegment(n))
}
