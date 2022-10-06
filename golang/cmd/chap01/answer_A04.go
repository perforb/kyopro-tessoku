package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	for i := 9; i >= 0; i-- {
		p := int(math.Pow(float64(2), float64(1)))
		fmt.Print(n / p % 2)
	}
}
