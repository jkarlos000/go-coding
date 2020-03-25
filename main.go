package main

import "fmt"

// La estructura Animal tiene un Nombre, una Edad, para representar un animal.
type Animal struct {
	Name string
	Age uint
}

// Función String() permite que el 'Animal' satisface una interfaz Stringer o conocida como @toString()
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	myStr := a.String()
	fmt.Println(myStr)
}