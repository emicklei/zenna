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
