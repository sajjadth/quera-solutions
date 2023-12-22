package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func floorDivide(n, x int) int {
	return n / x
}

func countDivisibleNumbers(n, a, b, c, d int) int {
	result := floorDivide(n, a) + floorDivide(n, b) + floorDivide(n, c) + floorDivide(n, d)

	result -= floorDivide(n, lcm(a, b)) + floorDivide(n, lcm(a, c)) + floorDivide(n, lcm(a, d)) +
		floorDivide(n, lcm(b, c)) + floorDivide(n, lcm(b, d)) + floorDivide(n, lcm(c, d))

	result += floorDivide(n, lcm(a, lcm(b, c))) + floorDivide(n, lcm(a, lcm(b, d))) + floorDivide(n, lcm(a, lcm(c, d))) +
		floorDivide(n, lcm(b, lcm(c, d)))

	result -= floorDivide(n, lcm(a, lcm(b, lcm(c, d))))

	return result
}

func main() {
	var n, a, b, c, d int

	fmt.Scanln(&n, &a, &b, &c, &d)

	fmt.Println(countDivisibleNumbers(n, a, b, c, d))
}
