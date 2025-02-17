package main

import "fmt"

type my_struct struct {
	name     string
	age      int
	phone_no int64
}

func main() {
	var nischal my_struct
	nischal.name = "Nischal"
	nischal.age = 19
	nischal.phone_no = 9749777419
	fmt.Println(nischal.name)
	fmt.Println(nischal.age)
	fmt.Println(nischal.phone_no)
}
