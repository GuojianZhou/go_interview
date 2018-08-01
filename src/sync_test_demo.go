package main

import (
	"fmt"
	"sync"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("1 + i: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("2 + i: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
