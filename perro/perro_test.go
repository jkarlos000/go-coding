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
func BenchmarkEdad(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		Edad(i)
	}
}
