package main

import "fmt"

func main() {
	a := make([]string, 4, 5) // make사용하면 a는 슬라이스
	// a := []string{"a","b","c","d"}
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2] // ["a", ""]
	fmt.Println(a, len(a), cap(a))
	as[1] = "z" // ["a","z"]
	fmt.Println(a, len(a), cap(a))
	fmt.Printf("%x %x %x \n", &a[0], &as[0], &a[1])

	fmt.Println(a, len(a), cap(a))
}
