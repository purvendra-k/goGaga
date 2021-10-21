package main

import "fmt"

var Element = struct {
	name string
	age  int
}{
	name: "Naman",
	age:  90,
}

func main() {

	fmt.Println(Element)
	fmt.Printf("%T", Element)

}
