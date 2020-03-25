package main

import "fmt"

func main() {
	// Diccionarios (asi se conoce en Python) de Go lang (MAP)
	moons := make(map[string]string)

	// Asignamos los valores
	moons["Earth"] = "Moon"
	moons["Jupiter"] = "Europa"
	moons["Saturn"] = "Titan"

	fmt.Println(moons)

	delete(moons, "Saturn")
	fmt.Println(moons)
}
