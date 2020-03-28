package main

import "fmt"

type animal struct {
	name string
	species string
}

func (a *animal) Name() {
	fmt.Println(a.name)
}

func (a *animal) Species() {
	fmt.Printf("%s belongs to %s species\n", a.name, a.species)
}

type gopher struct {
	color string
	animal
}

func (g *gopher) Color() {
	fmt.Printf("%s is the color: %s\n", g.name, g.color)
}

func (g *gopher) Name() {
	fmt.Printf("I'm a %s but my real name is %s\n", g.name, "Jkarlos")
}

func New(a animal) {
	fmt.Println("I'm similar to inheritance, but I'm composition")
}

func main() {
	a := animal{
		name:    "Gopher",
		species: "Go Gopher",
	}
	g := &gopher{
		color:  "blue",
		animal: a,
	}
	g.Name()
	g.Species()
	g.Color()

	New(g)
}
