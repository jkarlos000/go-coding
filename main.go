package main

import "fmt"

func main() {
	galaxy := "M87"
	switch galaxy {
	case "Mliky Way":
		fmt.Printf("Gaalaxy name is 'Milky Way'\n")
	case "Andromeda":
		fmt.Printf("Galaxy name is 'Andromeda'\n")
	case "M87":
		fmt.Printf("Galaxy name is 'M87'\n")
	default:
		fmt.Printf("Don't exits a Galaxy")
	}
}
