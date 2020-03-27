package main

import (
	"fmt"
	"strings"
)

func getName(names ...string) string {
	return strings.Join(names, " ")
}

func main() {
	fmt.Println(getName("Claude", "Debussy"))
	fmt.Println(getName("Ravel"))
	fmt.Println(getName(""))
}
