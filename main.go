package main

import "fmt"

type Person struct {
	Name, LastName string
	Age int
}

type Human interface {
	getName()
	getLastName()
	getFullName()
	getAge()
	setFullName(name string, lastname string)
}

func (p Person) getName() {
	panic("implement me")
}

func (p Person) getLastName() {
	panic("implement me")
}

func (p Person) getFullName() {
	panic("implement me")
}

func (p Person) getAge() {
	panic("implement me")
}

func (p *Person) setFullName(name string, lastname string) {
	p.Name = name
	p.LastName = lastname
}

func main() {
	p := &Person{
		Name:     "Juan Carlos",
		LastName: "Ulloa Vasquez",
		Age:      25,
	}
	fmt.Printf("%+v", *p)
	p.setFullName("Jose Eduardo", "Ulloa Pam")
	fmt.Printf("%+v", *p)
}
