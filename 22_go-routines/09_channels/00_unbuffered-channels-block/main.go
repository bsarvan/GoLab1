package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			fmt.Println("A : ", i)
		}
	}()

	go func() {
		for {
			val, ok := <-c
			fmt.Println(ok)
			fmt.Println("B : ", val)
		}
	}()

	time.Sleep(time.Second)
}
