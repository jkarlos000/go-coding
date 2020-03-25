package main

import "fmt"

func main() {
	a, b := 1, 10
	if a < b {
		fmt.Printf("%d is less than %d\n", a,b)
	} else if a == b {
		fmt.Printf("%d is equal to %d\n", a, b)
	} else {
		fmt.Printf("%d is greater than %d\n", a, b)
	}
}
