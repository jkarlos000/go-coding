package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan int)

	// enviar
	go enviar(par, impar, salir)

	// recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizando")
}

func recibir(par <-chan int, impar <-chan int, salir <-chan int) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el canal par:", v)
		case v := <-impar:
			fmt.Println("Desde el canal impar:", v)
		case v := <-salir:
			fmt.Println("Desde el canal salir:", v)
			return
		}
	}
}

func enviar(par chan<- int, impar chan<- int, salir chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			par <- j
		} else {
			impar <- j
		}
	}
	/*close(par)
	close(impar)*/
	salir <- 0
}
