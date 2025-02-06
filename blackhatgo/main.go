package main

import "fmt"

func main() {
	x := 1
	if x == 1 {
		fmt.Printf("equal")
	} else if x != 1 {
		fmt.Println("not equal ")
	} else {
		fmt.Println("never")
	}
}
