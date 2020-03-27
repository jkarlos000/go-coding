package main

import "fmt"

func leftRotation(a []int, size int, rotation int) []int {
	var newArray []int
	for i := 0; i < rotation; i++{
		newArray = a[1:size]
		newArray = append(newArray, a[0])
		a = newArray
	}
	return a
}

func main() {
	a := []int{1,2,3,4,5,6}
	fmt.Printf("input: %+v\n", a)
	rotation := 4
	fmt.Printf("output: %+v\n", leftRotation(a, len(a), rotation))
}
