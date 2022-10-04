package main

import "fmt"

func main() {
	var a, b int
	found := false
	_, _ = fmt.Scan(&a, &b)
	for a <= b {
		if 100%a == 0 {
			found = true
		}
		a++
	}
	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
