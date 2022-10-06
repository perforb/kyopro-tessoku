package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	solution()
}

func solution() {
	var n int
	_, _ = fmt.Scan(&n)
	answer, base := 0, 1
	for n > 0 {
		answer += n % 10 * base
		n /= 10
		base <<= 1
	}

	fmt.Println(answer)
}

func anotherSolution() {
	var sum int
	var n string
	_, _ = fmt.Scanf("%s", &n)
	digits := strings.Split(n, "")
	for i, j := 0, len(digits)-1; i < len(digits); i, j = i+1, j-1 {
		number, _ := strconv.Atoi(digits[i])
		sum += int(math.Pow(float64(2), float64(j))) * number
	}
	fmt.Println(sum)
}
