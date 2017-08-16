package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	fmt.Println("num cpus:\t\t", runtime.NumCPU())
	fmt.Println("goroutines start:\t", runtime.NumGoroutine())

	go func() {
		fmt.Println("hello from first one")
		wg.Done()
	}()

	fmt.Println("num cpus:\t\t", runtime.NumCPU())
	fmt.Println("goroutines middle:\t", runtime.NumGoroutine())
	go func() {
		fmt.Println("hello from second one")
		wg.Done()
	}()

	fmt.Println("num cpus:\t\t", runtime.NumCPU())
	fmt.Println("goroutines end:\t\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("about to end")
}
