package main

import "fmt"

// recursividad
func fibor(n int) int {
	if n <= 1 {
		return n
	}
	return fibor(n-1) + fibor(n-2)
}

// iterativo
func fiboi() func() int {
	x, y := 0, 1
	return func() int {
		r := x
		x, y = y, x+y
		return r
	}
}

func main() {

	// Funciones anónimas tipo Closure
	n := 10
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", fibor(i))
	}
	fmt.Println()
	next_fibo := fiboi()
	for i := 0; i <= n; i++ {
		fmt.Printf("%d ", next_fibo())
	}
	fmt.Println()

}