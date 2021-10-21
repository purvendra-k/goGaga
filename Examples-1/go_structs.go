package main

import (
	"fmt"
)

type width int //user defined type
type Mobile struct {
	brand string
	model string
	price int
}

func (mob Mobile) display(brand string) string { //func associated with Mobile
	fmt.Printf("Address inside display:- %p\n", &mob)
	mob.brand = brand
	fmt.Println("Inside display", mob.brand)
	fmt.Println(mob)
	fmt.Println("*******")
	return mob.brand
}
func (mob *Mobile) show(brand string) string { //func associated with Mobile
	fmt.Printf("Address inside show:- %p\n", mob)
	mob.brand = brand
	return mob.brand
}
func main() {
	var height width
	fmt.Println(height)
	m := Mobile{}
	fmt.Println(m) //Default values inside struct{" " 0}
	var mob Mobile //Instance creation using var
	fmt.Println(mob)
	mobs := new(Mobile)
	fmt.Println(mobs)
	phone := Mobile{"Samsung", "Galaxy", 24500} //Struct initialization
	fmt.Println("Before Change:", phone)
	fmt.Printf("%T\n", phone)
	fmt.Printf("Address inside main:- %p\n", &phone)
	fmt.Println("Function Call", phone.display("Xiaomi")) // Xiomi
	fmt.Println("After Change:", phone)                   //still old values will come {"Samsung","Galaxy",24500}
	fmt.Println("Function Call:", phone.show("Xiaomi"))   //calling show()
	fmt.Printf("Address inside main 2:- %p\n", &phone)
	fmt.Println("After Change:", phone) //Here changed values will reflect
}
