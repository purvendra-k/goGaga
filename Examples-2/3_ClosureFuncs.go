package main

import "fmt"

func main() {
	l := 20
	b := 30
	fmt.Printf("%p\n", &l)
	fmt.Printf("%p\n", &b)

	//Closures are a special case of anonymous functions.
	//Closures are anonymous functions which access the variables defined outside the body of the function.
	func(a int, c int) {
		l = 40
		b = 50
		fmt.Printf("%p\n", &l)
		fmt.Printf("%p\n", &b)
		var area int
		area = l * b
		fmt.Println(area)
		fmt.Println(a * c)
	}(l, b)

	fmt.Printf("%d, %p\n", l, &l)
	fmt.Printf("%d, %p\n", b, &b)
}
