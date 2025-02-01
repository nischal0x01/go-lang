package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello world")
	http.ListenAndServe(":8080", nil)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world")
	})
}
