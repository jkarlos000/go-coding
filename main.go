package main

import "fmt"

type Prueba struct {
	msg string
	id int
}

func main() {
	var value interface{}
	show(value)

	value = 49
	show(value)

	value = "empty interface"
	show(value)
	//value = new(Prueba)
	value = Prueba{
		msg: "Hello from World",
		id:  577,
	}
	show(value)
}

func show(value interface{}) {
	fmt.Printf("(%v, %T)\n", value, value)
}