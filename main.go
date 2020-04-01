package main

import "fmt"

func main() {
	//Canales unilaterales

	// Declarando canal de tipo SET ONLY, solo se envia información a este Canal, este tendrá un buffer de 2
	ca := make(chan<- int, 2)

	/*ca <- 42
	ca <- 43*/

	fmt.Printf("%T\n", ca)

	// Declarando canal de tipo RECEIVE ONLY, solo recibe información de este Canal, este tendrá un buffer de 2
	cb := make(<-chan int, 2)

	/*cb <- 42
	cb <- 43*/

	fmt.Printf("%T\n", cb)

	// Canales bidireccionales pueden ser HEREDADOS a los canales unilaterales receive or send, esta es la única operación permitida
	 c := make(chan int, 2)

	 ca = c
	 cb = c

	fmt.Printf("%T\n", c)
}
