package main

import "fmt"

func main() {

	a, b := 87, 20

	// we are using logical operators here.
	if a != b && a%2 == 0 {
		fmt.Println("Some First condition here")
	} else if a%3 == 0 {
		fmt.Println("Some Second condition here")
	} else {
		fmt.Println("Some Third condition here")
	}

	// assignment operators
	a += 20
	fmt.Println(a)

	b -= 3
	fmt.Println(b)

	// bitwise operators.

	fmt.Println(a & b)
}
