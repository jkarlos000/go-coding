package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre, Apellido	string
	Frases	[]string
}

func main() {
	p1 := persona{
		Nombre:   "Jame",
		Apellido: "Bond",
		Frases:   []string{"Shaken, not stirred", "¿Algún último deseo?", "Nunca digas nunca"},
	}

	buf, err := aJSON(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func aJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Hubo un error al convertir en aJSON: %v", err)
	}
	return bs, nil
}
