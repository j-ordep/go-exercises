package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan int) <-chan int {
	c := make(chan int)

	go func ()  {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()
	
	return c
}

func receive(c, q <-chan int) {
	for {
		select {
		case <-q:
			return
		case v := <- c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}