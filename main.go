package main

import (
	"fmt"
)

func greetings(name string) string {
	return "Hello " + name
}

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(greetings("Einstein"))
	fmt.Println(add(54, 56))
}
