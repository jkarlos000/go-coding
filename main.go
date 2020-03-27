package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan <- int)  {
	for j := range jobs {
		fmt.Println("worker", id, "inicio el trabajo:",j)
		time.Sleep(1*time.Millisecond)
		fmt.Println("worker",id,"termino el trabajo",j)
		results <- j*2
	}
}

func main() {
	// Dos canales - se envía el trabajo y se recolecta el resultado
	jobs := make(chan int, 10)
	results := make(chan int, 100)

	// Iniciamos primero con 3 trabajadores, inicialmente bloqueados debido a que no hay trabajos hechos
	for w:=1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Aquí enviaremos 5 'trabajos' y también 'cerraremos' el canal indicando que todos los trabajos están hechos
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	//time.Sleep(200*time.Millisecond)
	// recolectamos todos los resultados del trabajo
	for a:=1; a<=5; a++ {
		<-results
	}
}
