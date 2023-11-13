package main

import "fmt"

func main() {
	a := make([]string, 4, 5) // make사용하면 a는 슬라이스
	// a := []string{"a","b","c","d"}
	a[0] = "a"
	a[2] = "c"
	a[3] = "d"
	as := a[0:2] // ["a", ""]
	// c := append(a, "y") // 메모리공간이 늘어나지 않는다
	c := append(a, "y", "x") // 메모리공간을 늘린다.
	fmt.Println(a, len(a), cap(a))
	fmt.Println(a, len(c), cap(c))
	as[1] = "z" // ["a","z"]
	// 	fmt.Println(a, len(a), cap(a))
	// 	fmt.Printf("%x %x %x \n", &a[0], &as[0], &a[1])

	// 	fmt.Println(a, len(a), cap(a))
}
