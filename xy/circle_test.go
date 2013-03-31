package xy

import "fmt"

func ExampleCircle() {
	c := Circle{Point{0, 0}, 1, 0, 0}
	fmt.Println(c)
	// Output:
	// Circ( P(0,0), 1 )
}
