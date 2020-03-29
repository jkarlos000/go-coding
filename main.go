package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0
	for {
		x ++
		if x%2 != 0 {
			fmt.Println("Hehehehe, sopenco, tu sos",x)
			//fmt.Printf("Hehehehe, sopenco, tu sos %v", x)
			continue
		}
		fmt.Println(x)
		time.Sleep(time.Second*1)
	}
}
