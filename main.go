package main

import (
	"fmt"
	"math"
)

type point struct {
	x,y,z int
}

func main() {
	p := point{
		x: 10,
		y: 200,
		z: 3000,
	}
	fmt.Println(p)
	fmt.Printf("%v\n",p)
	fmt.Printf("%+v\n",p)
	fmt.Printf("%#v\n",p)
	fmt.Printf("%T\n",p)
	fmt.Println()

	var b bool
	b = true
	fmt.Printf("%t\n",b)
	fmt.Println()

	fmt.Printf("%d\n", 12335)
	fmt.Println()

	fmt.Printf("%b\n", []byte("HOLA MUNDO"))
	fmt.Printf("%b\n", 9)
	fmt.Println()

	fmt.Printf("%c\n", 97)
	fmt.Println()

	fmt.Printf("%x\n", 1501)
	fmt.Println()

	fmt.Printf("%f\n", 3.14159268)
	fmt.Println()

	fmt.Printf("%e\n", math.Pow(-9,3))
	fmt.Printf("%E\n", math.Pow(-5,15))
	fmt.Println()

	fmt.Printf("%s\n","String Format")
	fmt.Printf("%s\n", "comillas")
	fmt.Printf("%s\n", "\"comillas\"")
	fmt.Printf("%q\n", "\"comillas\"")
	fmt.Println()

	fmt.Printf("%x\n", "Hexadecimeame esto")
	fmt.Printf("%x\n", "h")
	fmt.Printf("%x\n", "e")
	fmt.Printf("%x\n", "x")
	fmt.Printf("%x\n", " ")
	fmt.Println()

	fmt.Printf("%p\n", &p)
	fmt.Println()
}
