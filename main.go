package main

import "fmt"

type persona struct {
	nombre	string
	apellido	string
}

type agente struct {
	persona
	administrador bool
}

type humano interface {
	presentar()
}

func (p persona) presentar()  {
	fmt.Println(p)
}

func (a agente) presentar() {
	fmt.Println(a)
}

func foo(h humano) {
	switch h.(type) {
	case agente:
		fmt.Println("Soy un agente, mi nombre es: ", h.(agente).nombre)
	case persona:
		fmt.Println("Soy un hombre ordinario, una persona simple, mi nombre es: ", h.(persona).nombre)
	default:
		fmt.Println("No pertenezco a la estructura, pero si a la interfaz.")
	}
}

func main() {
	p := persona{
		nombre:   "Widito",
		apellido: "Bueno",
	}
	a := agente{
		persona:       persona{
			nombre:   "Pi, PiBond",
			apellido: "00Pi",
		},
		administrador: true,
	}
	foo(p)
	foo(a)
}
