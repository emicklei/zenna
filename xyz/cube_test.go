package xyz

import "fmt"

func ExampleCube() {
	c := Cube(P(1, 2, 3), P(4, 5, 6))
	fmt.Println(c)
	fmt.Println(c.Origin)
	fmt.Println(c.Corner())
	// Output:
	// Cube(P(1,2,3), P(4,5,6))
	// P(1,2,3)
	// P(5,7,9)
}
