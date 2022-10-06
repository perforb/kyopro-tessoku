package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for i := 9; i >= 0; i-- {
		p := powInt(2, i)
		fmt.Print(n / p % 2)
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
