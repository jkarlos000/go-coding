package perro

import (
	"fmt"
	"testing"
)

func TestEdad(t *testing.T) {
	v := Edad(15)
	if v != 105 {
		t.Error("Expect", 105, "Got", v)
	}
}

func ExampleEdad() {
	fmt.Println(Edad(15))
	//Output:
	//105
}