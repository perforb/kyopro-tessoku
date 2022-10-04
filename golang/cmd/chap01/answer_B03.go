package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	m := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&m[i])
	}
	found := false

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if m[i]+m[j]+m[k] == 1000 {
					found = true
					break
				}
			}
		}
	}

	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
