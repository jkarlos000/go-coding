package main

import "fmt"

// retornamos una función y que esta retorna un entero
func init_seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Funciones anónimas tipo Closure
	next_int := init_seq()
	fmt.Println(next_int())
	fmt.Println(next_int())
	fmt.Println(next_int())

	next_int2 := init_seq()
	fmt.Println(next_int2())
	fmt.Println(next_int2())
	fmt.Println(next_int2())

}