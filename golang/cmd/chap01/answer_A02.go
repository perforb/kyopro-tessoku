package main

import (
	"fmt"
)

func main() {
	var n, x int
	var a int
	found := false
	_, _ = fmt.Scan(&n, &x)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&a)
		if a == x {
			found = true
			break
		}
	}
	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
