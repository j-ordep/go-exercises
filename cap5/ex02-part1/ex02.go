package main

import "fmt"


func main() {

	a := (10 == 10)
	b := (10 != 10)
	c := (10 <= 100)
	d := (10 < 10)
	e := (10 >= 100)
	f := (10 > 10)

	fmt.Printf("%v \n%v \n%v \n%v \n%v \n%v\n", a,b,c,d,e,f)

}