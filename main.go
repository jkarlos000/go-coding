package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{
		Name: "Interface",
		Body: "Empty Interface",
		Time: 15567476233,
	}
	b, err := json.Marshal(m)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("b = %T%+v\n", b, b)
	fmt.Printf("b = %T%s", b, b)
}
