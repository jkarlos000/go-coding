package main

import "fmt"

func main() {
	var palabras []int
	//palabras = make([]int,0,10)
	//palabras = append(palabras, 15,39,45,100)
	palabras = append(palabras, 10000, 495, 193)
	fmt.Printf("Palabras: longitud %d, capacidad %d, data: %v\n", len(palabras), cap(palabras), palabras)
}
