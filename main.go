package main

import "fmt"

func main() {
	/*a := 10
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T %T", a, b)*/

	i, j := 42, 2701

	p := &i // Definimos a 'p' como puntero a 'i'
	fmt.Println(*p) //Leemos 'i' como un puntero obteniendo el valor
	*p = 21 // Establecemos el valor de 'i', con nuestro puntero, recuerda que estamos manipulando memoria con el puntero, para obtener el valor usamos *
	fmt.Println(i) // Obtenemos el nuevo valor de 'i'

	p = &j // Apuntamos a 'j'
	*p = *p / 37 // Dividimos el valor de 'j' mediante el puntero
	fmt.Println(j) //Miramos el nuevo valor de 'j'
}