package main

import "fmt"

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7 ])
	fmt.Println(slice[2:9])

	fmt.Println(slice[2:len(slice)-1])

}