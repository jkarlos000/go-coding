package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(8)
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Verificando el error 1:", ctx.Err())
	fmt.Println("Número de Go Runtime 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <- ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Trabajo muy duro, como un esclavo, denme dinero +", n)
			}
		}
	}()
	time.Sleep(time.Second*2)

	fmt.Println("Verificando error:", ctx.Err())
	fmt.Println("Número de GORuntime 2da:", runtime.NumGoroutine())

	fmt.Println("A punto de cancelar el context")

	cancel()

	fmt.Println("Context cancelado")

	time.Sleep(time.Second*2)
	fmt.Println("Verificando error 3:", ctx.Err())
	fmt.Println("Número de GORuntime 3ra:", runtime.NumGoroutine())

}
