package main

import (
	"fmt"
	"io"
)

type gopher struct {

}

func (g *gopher) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (g *gopher) Read(p []byte) (n int, err error) {
	return 0, nil
}

func fetchType(i interface{}) {
	switch i.(type) {
	case io.ReadWriter:
		fmt.Println("Implementado ReadWriter Interface")
	default:
		fmt.Println("No tengo implementado la interfaz ReadWriter")
	}
}

func main() {
	g := &gopher{}
	fetchType(g)
}
