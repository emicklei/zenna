package xy

import (
	"fmt"
	"math"
)

func ExamplePoint_RotatedBy() {
	v := Point{1, 0}
	r := v.RotatedBy(math.Pi / 4)
	fmt.Println(r)
	// Output:
	// P(0.7071067811865476,0.7071067811865475)
}

func ExamplePoint_Polar() {
	v := Polar(10, math.Pi/4)
	fmt.Println(v)
	// Output:
	// P(7.0710678118654755,7.071067811865475)
}
