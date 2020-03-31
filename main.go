package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `Cambiamejk1!.`
	s2 := "Hola"
	bs1, err := bcrypt.GenerateFromPassword([]byte(s), 4)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs1))
	var bs []byte
	bs, err = bcrypt.GenerateFromPassword([]byte(s2), 4)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	err = bcrypt.CompareHashAndPassword(bs1, []byte("Cambiamejk1!."))
	if err != nil {
		fmt.Println("No te puedes loggear")
		return
	}
	fmt.Println("Te has loggeado")
}
