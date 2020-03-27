package main

import (
	"fmt"
	"time"
)

func sleep_250()  {
	fmt.Println("sleep_250")
	time.Sleep(250*time.Millisecond)
}

func sleep_500() {
	fmt.Println("sleep_500")
	time.Sleep(500*time.Millisecond)
}

func main() {
	go sleep_250()
	go sleep_500()
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Main")
}
