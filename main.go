package main

import "fmt"

const (
	a = 42
	b int = 43
)

func main() {
	x := 42
	fmt.Printf("%d\t\t%b\t\t%#x\n",x,x,x)

	fmt.Println(42 == 42)
	fmt.Println(43 <= 43)
	fmt.Println(42>=41)
	fmt.Println(42 != 43)
	fmt.Println(42 < 43)
	fmt.Println(42>41)

	fmt.Println(a,b)

	y := 42
	fmt.Printf("%d\t%b\t%#x\n", y,y,y)
	z := y << 1
	z = z | 1
	fmt.Printf("%d\t%b\t%#x\n", z,z,z)

}
