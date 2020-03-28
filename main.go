package main

import "fmt"

type gopher struct {
	Name	string
	Color	string
	Size	float64
}

func main() {
	//var gopher gopher
	var gopherP *gopher
	//fmt.Println(gopher == nil)
	fmt.Println(gopherP == nil)
	fmt.Println("Hello, Playground")


	var integer *int
	var empty interface{}

	fmt.Println(integer == nil)
	fmt.Println(empty == nil)
	fmt.Println(integer == empty)

	empty = integer

	fmt.Println(integer == nil)   // Prints: true
	fmt.Println(empty == nil)     // Prints: false
	fmt.Println(integer == empty) // Prints: true
	fmt.Printf("a=(%T,%v)\n", empty, empty) // Prints: a=(*int,<nil>)
}
