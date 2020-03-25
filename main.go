package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Método (recibe un valor)
func (p Person) hello() string {
	return "Hello, I'm " + p.firstName + " " + p.lastName + ", " + strconv.Itoa(p.age) + " years old"
}

// Método (recibe un puntero) - Con esto se modifica los datos de la estructura
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
//inicializando la persona como un 'struct'
	p1 := Person{
		firstName: "Steven",
		lastName:  "King",
		city:      "Chicago",
		gender:    "m",
		age:       23,
	}
	p2 := Person{"Neena", "Kochhar", "Boston", "f", 13	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p2.firstName, p2.lastName)
/*	p2.age++
	fmt.Println(p2)*/
	fmt.Println(p1.hello())
	fmt.Println(p2.hello())

	p2.hasBirthday() // Edad + 1
	fmt.Println(p2.hello())
	p2.hasBirthday() // Edad + 1
	fmt.Println(p2.hello())
}