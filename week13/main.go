package main

import "fmt"

func main() {
	var a []string
	a = make([]string, 4, 5)

	fmt.Printf("%#v\n", a) // zero value of empty slice

	fmt.Println(a, len(a), cap(a))
}
