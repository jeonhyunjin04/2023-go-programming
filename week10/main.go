package main

import "week10/src/greeting"
	   "week10/src/mymath"

func main() {
	greeting.Hello()
	fmt.Println(mymath.MyAbs(99))
	fmt.Println(mymath.MyAbs(-7))
	greeting.Hi()
}
