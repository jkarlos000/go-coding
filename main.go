package main

import "fmt"

type Chichis struct {
	next *Chichis
	apodo string
	field string
	edad int
}

type ChichitosList struct {
	head *Chichis
	name string
}

func createChichitosList(n string) *ChichitosList {
	return &ChichitosList{
		name: n,
	}
}

func (l_list *ChichitosList) addChichis(a string, f string, y int) error {
	fmt.Printf("Agregando Chichito %s %s %d\n", a,f,y)
	l := &Chichis{
		apodo: a,
		field: f,
		edad:  y,
	}
	if l_list.head == nil {
		l_list.head = l
	} else {
		currentNode := l_list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = l
	}
	return nil
}

func (l_list *ChichitosList) deleteChichis(a string, f string, y int) error {
	fmt.Printf("Eliminando Chichito: %s %s %d\n", a, f, y)

	currentNode := l_list.head
	if currentNode == nil {
		fmt.Println("No hay Chichitos!!!! ayuda!")
		return nil
	}

	// head
	if currentNode.apodo == a && currentNode.field == f && currentNode.edad == y {
		if currentNode == l_list.head {
			l_list.head = currentNode.next
		}
		return nil
	}

	// others

	fmt.Printf("*currentNode %+v\n", *currentNode)
	for currentNode.next != nil  {
		fmt.Printf("*currentNode.next: %+v\n", *currentNode.next)
		next := currentNode.next
		if next.apodo == a && next.field == f && next.edad == y {
			fmt.Println("El chichito buscado fue encontrado en el enlace siguiente!")
			currentNode.next = next.next
			break
		}
		currentNode = currentNode.next
	}
	return nil
}

func (l_list *ChichitosList) showAllChichitos() error {
	fmt.Printf("\nLista actual de Chichitos:\n")
	fmt.Println("----------------------")
	currentNode := l_list.head
	if currentNode == nil {
		fmt.Println("Lista de Chichitos Vacía")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil  {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
	fmt.Println("----------------------")
	return nil
}

func main() {
	myChichitos := createChichitosList("MyListaDeChichis")
	fmt.Println("Se ha creado una lista vacia de Chichitos!")

	myChichitos.addChichis("BB", "El bebito", 20)
	myChichitos.addChichis("Bichito", "El bichito de Luz", 16)
	myChichitos.addChichis("Chichiiii", "Chiichiiiiiii", 10)
	myChichitos.addChichis("Sergio", "Sergio Mauricio", 0)
	myChichitos.showAllChichitos()

	myChichitos.deleteChichis("Sergio", "Sergio Mauricio", 0)
	myChichitos.deleteChichis("Bichito", "El bichito de Luz", 16)
	myChichitos.showAllChichitos()

	myChichitos.addChichis("Gordo", "BB Gordo y Feo", 20)
	myChichitos.showAllChichitos()
}
