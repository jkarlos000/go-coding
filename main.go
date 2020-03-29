package main

import (
	"fmt"
)

func main() {
	s1 := `El BB:
Buenos dias my bb, tengo algo que mostrar:

*El bb no sale afuera
*El bebe necesita su leche.


Atte. Jk.`

	fmt.Println(s1)
	bs := []rune(s1)

	fmt.Printf("%T\n", bs)
	fmt.Printf("%v\n", bs)

	fmt.Println()

	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i])
	}

	fmt.Println()

	for i, v := range s1 {
		fmt.Printf("El índice %d tiene el valor %v (%c) Hexa(%#x)\n", i, v,v,v)
	}
}
