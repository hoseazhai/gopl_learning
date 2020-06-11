package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func wordFrea(words string) {
	counts := make(map[string]int)

	input := bufio.NewScanner(strings.NewReader(words))
	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%v \t %v \n", k, v)
	}
}

func main() {
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%v \t %v \n", k, v)
	}
}
