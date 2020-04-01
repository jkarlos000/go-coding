package main

import "fmt"

func main() {

	// Canal sin buffer, esto requiere que trabajen dos runtime, en este caso main y la función literal
	ca := make(chan int)

	go func() {
		ca <- 42
	}()

	fmt.Println(<-ca)

	// Canal con buffer, funciona en la misma subrutina, [1] es el tamaño del buffer
	c1 := make(chan int, 1)

	c1 <- 42
	//c1 <- 43 // El tamaño de nuestro buffer es 1, si quiero añadir otro valor esto dara deadlock

	fmt.Println(<-c1)

	// optimización para multiples tamaños
	c2 := make(chan int, 2)

	c2 <- 42
	c2 <- 43

	fmt.Println(<-c2)
	fmt.Println(<-c2)

}
