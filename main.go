package main

import "fmt"

func main() {
	/*for _, i := range []int{1,2,3,4,5}{
		fmt.Println(i)
	}*/
	ids := []int{0,10,20,30,40,50,60}
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Si usted no tiene pensado usar el indexx
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// sumatoria de los ID
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)
}
