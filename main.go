package main

import "fmt"

func main() {
	c := make(chan int)
	//Enviar
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	//Recibir
	for v := range c{
		fmt.Println(v)
	}

	fmt.Println("Finalizado")

}
