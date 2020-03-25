package main

import "fmt"

func main() {
	// range usando string, nos devuelve el index y rune
	/*for i, c := range "CHICHITOMANCO" {
		fmt.Printf("%#U starts at byte position %d\n", c, i)
	}*/
	// range usando MAPS
	/*moons := map[string]string{"Earth":"Moon", "Jupiter":"Europa","Saturn":"Titan"}
	for k, v := range moons {
		fmt.Printf("%s: %s\n", k, v)
	}*/
	/*numbers := map[string]int{
		"Uno": 1,
		"Dos": 2,
		"Tres": 3,
		"Cuatro": 4,
		"Cinco": 5,
	}
	for k,v := range numbers {
		fmt.Println(k, v)
	}*/
	// Iteraremos sobre 5 valores en el canal 'queue'
	/*queue := make(chan string, 5)
	queue <- "Enceladus"
	queue <- "Titan"
	queue <- "Europa"
	queue <- "Ganemede"
	queue <- "Io"
	close(queue)
	// Este 'rango' itera sobre los elementos que son recibidos desde 'queue'. Debido a que hemos 'closed' the canal tambien, la interacion termino despues de recibir 5 colas
	for q := range queue {
		fmt.Println(q)
	}*/
	// Creamos un Slice vacio apuntado a nuestra estructura
	cities := []*Cities{}

	// Creamos una estructura y la agregamos al slice
	ct := new (Cities)
	ct.name = "London"
	ct.location[0] = 5
	ct.location[1] = 0
	cities = append(cities, ct)

	//Creamos otra estructura
	ct = new(Cities)
	ct.name = "Sydney"
	ct.location = [2]int{34, 51}
	cities = append(cities, ct)

	for i := range(cities){
		c := cities[i]
		fmt.Println("City:",*c)
	}
}

type Cities struct {
	name string
	location [2]int
}