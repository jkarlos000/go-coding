package main

import "fmt"

func fibonacci(c, quit chan int) {
	fmt.Println("Serie Fibonacci")
	x, y := 0, 1
	for {
		fmt.Printf("x = %d\n",x)
		select {
		case c <- x:
			x, y = y, x+y
		case quit_value := <-quit:
			fmt.Printf("quit_value = %d\n", quit_value)
			return
		default:
			fmt.Println("Sin actividad")
		}
	}
}

func main() {
	c := make(chan int)
	q := make(chan int)
	go func() {
		fmt.Println("GoRoutine Iniciando")
		for i := 0; i < 10; i++ {
			value := <-c
			fmt.Printf("received %d\n", value)
		}
		q <- 999
	}()
	fmt.Println("Ejecutando fibonacci()")
	fibonacci(c,q)
}
