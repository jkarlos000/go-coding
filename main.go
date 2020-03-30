package main

import "fmt"

func main() {
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}

	m["Carlos"] = 25
	for key, value := range m {
		fmt.Println(key, value)
	}
	delete(m, "Carlos")
	for key, value := range m {
		fmt.Println(key, value)
	}
	if v, ok := m["Carlos"]; ok {
		fmt.Printf("Existe la llave valor Carlos con edad %v\n", v)
		delete(m, "Robin")
	}
}
