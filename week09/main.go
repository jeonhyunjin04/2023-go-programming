package main

import "fmt"

// pass by pointer (call by pointer)
func swap(n1 *int, n2 *int) { // 두 수 교환
	temp := *n1
	*n1 = *n2
	*n2 = temp
}

func main() {
	a := 10
	b := 20

	c := &a               // integer 타입의 변수 주소는 integer 타입의 포인터 변수로만 받을 수 있다
	fmt.Printf("%T\n", c) // *int

	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)

}

//package main

//import "fmt"

//func double(number *int) {
//	*number = *number * 2
//}

//func main() {
//	var amount int = 5
//	double(&amount)
//	fmt.Printf("%d\n", amount)
//}

//package main

//import "fmt"

//func double(number int) {
//	number = number * 2
//}

//func main() {
//	var ammount int = 5
//	double(ammount)
//	fmt.Printf("%d\n", ammount)
//}

//package main

//import "fmt"

//func main() {
// 포인터 : 메모리 주소를 값으로 가지는 타입. & 주소연산자, * 포인터 기호
//	var a int = 10
//	var pa *int
//	pa = &a

//	fmt.Println(&a, a)
//	fmt.Println(pa, *pa)
//	fmt.Println(&pa)
//	fmt.Printf("%T %T\n", a, pa)
//}
