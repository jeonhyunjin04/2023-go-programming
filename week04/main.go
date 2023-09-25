package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var a int = 9
	var b = 2.7
	var c, d string
	var e float32 = 3.14
	var f string
	var g bool
	c = "inha~"
	//b = 2.7
	//a = 9
	d = "Go..."
	h := 'Z'
	i := "문자열"
	J := "변수명이 대문자로 시작하면 다른 패키지에서 이 변수를 사용할 수 있음"
	koreanZombie := "정찬성"

	fmt.Println(float64(a) > b)
	fmt.Println(a * int(b))

	fmt.Println(reflect.TypeOf(b))
	fmt.Println(a, b, c, d, e, f, g, h, i, J, koreanZombie)
	fmt.Println(reflect.TypeOf(h))
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(100))
	fmt.Println(reflect.TypeOf(2.71))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf("Go!"))
	// fmt.Println(math.Floor("inha"))
	// fmt.Println(strings.Title(3.14))
	fmt.Println('A')       // rune (unicode)
	fmt.Println('김', '\n') // rune (unicode)
	fmt.Println(math.Floor(2.15))
	fmt.Println(math.Ceil(2.15))
	fmt.Println(strings.Title("오픈소스\t프로그래밍~\n\"Go\""))
	fmt.Println(strings.Title("open source programming go!"))
}
