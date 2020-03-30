package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5}
	y := make([]int,0)
	y = append(y, 55,221,345,4848)
	x = append(x, y...)

	fmt.Println(x)
	x = append(x[:2],x[4:]...)

	fmt.Println(x)

}
