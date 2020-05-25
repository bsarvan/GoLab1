package main

import (
	"fmt"
	"sync"
)

func main() {

	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		fmt.Println("Done executing go routine A")
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		fmt.Println("Done executing go routine B")
		wg.Done()
	}()

	go func() {
		fmt.Println("Waiting on the go routines for completion")
		wg.Wait()
		fmt.Println("Closing the channel")
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
