package main

import "fmt"

func suma(xi ...int) int {
	total := 0
	for _, v := range xi{
		total += v
	}
	return total
}

func sumaPar(f func(ii ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi{
		if v % 2 == 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}

func main() {
	s1 := []int{22,13,54,23,6,3,72,76,2,4,12}
	x := sumaPar(suma,s1...)
	fmt.Println(x)
}
