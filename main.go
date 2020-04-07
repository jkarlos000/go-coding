package main

import "fmt"

type errorPer struct {
	info	string
}

func (ep errorPer) Error() string {
	return fmt.Sprintf("El error es: %v", ep.info)
}

func main() {
	e1 := errorPer{info: "Necesito dormir más"}
	foo(e1)
}

func foo(e error) {
	fmt.Println("foo corrió -", e, "\n", e.(errorPer).info) // Uso de asertion, es diferente a conversion
}
