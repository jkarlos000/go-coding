package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Número de CPU:", runtime.NumCPU())
	fmt.Println("Número de Go-Runtime:", runtime.NumGoroutine())
	runtime.GOMAXPROCS(8) //Cantidad de CPU a usar (mi maquina soporta 8)
	var contador int64

	const gs = 99999
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Número de Go-Runtime:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Cuenta:", contador)
}
