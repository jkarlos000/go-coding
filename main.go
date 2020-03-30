package main

import "fmt"

func main() {
	for i := 0; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%#U\n",i,i,i)
		switch {
		case i%2 == 0, i > 120://Cualquiera de estas condiciones sea verdadero, ejecutara la instrucción
			fmt.Println(i)
		case i == 15:
			fmt.Println("Soy 15")
			fallthrough
		case i == 2:
			fmt.Println("SOy 2 y 15 luego")
		default:
			fmt.Println("Default")
		}
	}
}
