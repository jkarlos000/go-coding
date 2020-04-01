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

// Uso del idioma Coma, el canal devuelve 2 parámetros, values, estado (false -> cerrado por close(), true -> funcional)
func recibir(par <-chan int, impar <-chan int, salir <-chan int) {
	for {
		select {
		case v := <-par:
			fmt.Println("Desde el canal par:", v)
		case v := <-impar:
			fmt.Println("Desde el canal impar:", v)
		case v, ok := <-salir:
			if !ok { // Si es falso, el canal se encuentra cerrado
				fmt.Println("Desde el canal salir:", v)
				return
			} else { // El canal aun se encuentra abierto.
				fmt.Println("Desde el canal salir:", v)
			}
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
	close(salir)
}
