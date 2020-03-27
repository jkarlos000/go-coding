package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("xd/defer.txt")
	if err != nil {
		panic("No se puede crear el fichero")
	} else {
		fmt.Printf("err = %+v", err)
	}
	defer f.Close()
	fmt.Fprint(f,"Hahaha soy parte de un fichero")
}
