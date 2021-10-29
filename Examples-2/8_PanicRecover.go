package main

import "fmt"

func main() {
	panicFunc()
	fmt.Println("Hello")
}

func panicFunc() {
	var action int
	// An anonymous function
	defer func() {
		action := recover()
		fmt.Println("Hello from defer")
		fmt.Println(action)
	}()

	fmt.Println("Enter 1 for Student and 2 for Professional")
	fmt.Scanln(&action)
	/*  Use of Switch Case in Golang */
	switch action {
	case 1:
		fmt.Printf("I am a  Student")
	case 2:
		//Generating Panic
		var i interface{}
		i = "Hello"
		s := i.(int)
		fmt.Println(s)
	default:
		panic(fmt.Sprintf("I am a  %d", action))
	}
}
