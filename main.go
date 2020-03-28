package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name, Body string
	Time int64
}

func main() {
	m := Message{
		Name: "Interfaz",
		Body: "Interfaz Vacía",
		Time: 1556747623,
	}
	b, err := json.Marshal(m)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("b = %T%+v\n", b, b)
	fmt.Printf("b = %T%s\n",b,b)

	var md Message
	err = json.Unmarshal(b, &md)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("b = %T%+v\n", b, b)
	fmt.Printf("b = %T%s\n",b,b)
	fmt.Printf("type = %T, md = %+v\n", md, md)
}
