package main

import "fmt"

func foo() {
	fmt.Println("Foo()")
	defer bar()
	panic("Argh!")
}

func bar() {
	fmt.Println("bar()")
	err := recover()
	fmt.Printf("err = %+v", err)
}

func main() {
	foo()
	fmt.Println()
}
