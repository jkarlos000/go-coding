package main

import "fmt"

type Cities struct {
	name string
	location [2]int
}
func main() {
	// Creamos un Slice vacío, apuntando a la estructura
	cities := []*Cities{}

	// Creamos una estructura y la agregamos a nuestro Slice
	ct := new (Cities)
	ct.name = "London"
	ct.location = [2]int{5,0}
	cities = append(cities, ct)

	// Creamos otra estructura de tipo Cities
	ct = new(Cities)
	ct.name = "Sydney"
	ct.location = [2]int{34, 51}
	cities = append(cities, ct)

	for i:= range (cities) {
		c := cities[i]
		fmt.Println("City:", *c)
	}
}