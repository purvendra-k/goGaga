package main

import (
	"fmt"
	"reflect"
)

func main() {
	variadicExample_1("red", "blue", "green", "yellow")

	variadicExampl_2()
	variadicExampl_2("red", "blue")
	variadicExampl_2("red", "blue", "green")
	variadicExampl_2("red", "blue", "green", "yellow")

	fmt.Println(area("Rectangle", 20, 30))
	fmt.Println(area("Square", 20))

	variadicExample_3(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})

	fmt.Println(myFunc(5))

	mySlice := make([]interface{}, 3, 5)
	mySlice[0] = 1
	mySlice[1] = true
	mySlice[2] = "Hello"

	// Comma Value Idiom
	for _, val := range mySlice {
		fmt.Printf("%v, %T\n", val, val)
	}

}

func variadicExample_1(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])
}

func variadicExampl_2(s ...string) {
	fmt.Println(s)
}

func area(str string, y ...int) int {

	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

func variadicExample_3(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind()) // See the use of reflection
	}
}

func myFunc(x int) interface{} {
	return "Hello"
}
