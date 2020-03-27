package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello from GoRoutine")
	fmt.Println("Hello from 1")
	fmt.Println("Hello from 2")
	fmt.Println("Hello from 3")
}
func main() {
	go hello()
	fmt.Println("Main")
	time.Sleep(1 * time.Millisecond)
}
