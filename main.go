package main

import "fmt"

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\n",a,a)
	b := a << 1
	fmt.Printf("%d\t\t%b\n",b,b)

	kb := 1024
	gb := kb * 1024
	tb := gb * 1024

	fmt.Printf("%d\t\t%b\n",kb,kb)
	fmt.Printf("%d\t\t%b\n",gb,gb)
	fmt.Printf("%d\t\t%b\n",tb,tb)
}
