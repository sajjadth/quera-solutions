package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func organizeMovieNames(movies []string) string {
	for i, movie := range movies {
		movies[i] = strings.Title(strings.ToLower(movie))
	}

	return strings.Join(movies, "\n")
}

func main() {
	var n int
	var movies []string

	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		movies = append(movies, s)
	}

	fmt.Println(organizeMovieNames(movies))

}
