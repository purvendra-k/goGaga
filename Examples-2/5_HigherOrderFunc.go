package main

/*
Function returning or accepting a function is called higher order function.
Simple functions are called first order functions.
*/

import "fmt"

func sum(x, y int) int {
	return x + y
}

func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}
func main() {
	partial := partialSum(3)
	fmt.Printf("%T\n", partial)
	fmt.Println(partial(7))
}
