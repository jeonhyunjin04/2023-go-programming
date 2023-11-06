package main

import "fmt"

func main() {
	primes := [3]int{2, 3, 5} //initialize by array literal
	//primes[2] = 6
	for i := 0; i < 3; i++ {
		fmt.Println(i, primes[i])
	}
	// 초기화 하지 않은 원소의 제로 값은 해당 원소 타입의 제로값으로 결정된다.
	test := [5]bool{true, true, true}
	fmt.Println(test[3])
	fmt.Println(test)
	//fmt.Println(test[5]) // 컴파일에러, invalid argument: index 5 out of bounds [0:5]

	i := 0
	//for i < 6 { // 런타임에러, panic : runtime error: index out of range [5] with length 5
	for i < len(test) {
		fmt.Println(test[i])
		i++
	}
}
