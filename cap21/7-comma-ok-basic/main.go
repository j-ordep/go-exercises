package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 42
		close(canal)
	}()

	v, ok := <-canal

	fmt.Println(v, ok) // 42, true

	v, ok = <-canal

	fmt.Println(v, ok) // 0, false
}