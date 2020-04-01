package main

import "fmt"

func main() {
	c := make(chan int)

	// enviar
	go enviar(c)
	/*go*/ recibir(c)

	fmt.Println("Finalizado")

}

func enviar(c chan<- int) {// Solo se recibe un canal SEND ONLY
	c <- 42
}

func recibir(c <-chan int) {// SOLO se recibe un canal RECEIVE ONLY
	fmt.Println(<-c)
}
