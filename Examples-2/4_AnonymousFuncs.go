package main

import "fmt"

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func main() {
	fmt.Println(area(20, 30))

	// Pass args to anonymous func
	func(l int, b int) {
		fmt.Println(l * b)
	}(30, 40)

	func(l int, b int) {
		fmt.Println(l * b)
	}(30, 40)

	//Function with a return value
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)

	// Anonymous function inside a for loop
	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}
}
