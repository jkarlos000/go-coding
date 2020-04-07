package main

import "fmt"

func main() {
	fmt.Println("la suma de 2 + 4 es:", miSuma(2,4))
	fmt.Println("la suma de 1 + 5 es:", miSuma(1,5))
	fmt.Println("la suma de 6 + 7 es:", miSuma(6,7))
}

func miSuma(xi ...int) int {
	var sum int
	for _,v := range xi {
		sum += v
	}
	return sum
}
