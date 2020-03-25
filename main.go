package main

import "fmt"

func say_hello(msg string) {
	fmt.Println(msg)
}

/*
Esta es una función regular
nombre: regular_f_returning_anonymous_f
return type: func(string) - Esta función anónima requiere como parámetro un string
*/
func regular_f_returning_anonymous_f() func(string) {
	// devuelve una función anónima que es una función interna
	return func(msg string) {
		fmt.Println(msg)
	}
}

func main() {
	//Funciones anónimas.... hay que practicar mucho esta parte

	// función regular
	//say_hello("Hello from regular function!")

	// función anónima
	func(msg string){
		fmt.Println(msg)
	}("Hello from an anonymous function")

	print_fnc := regular_f_returning_anonymous_f()
	print_fnc("Hello from returned anonymous function")
	// Todavía no se que uso darle a esto, pero es necesario investigar y mucho, sobre todo para las promesas

}