package xy

import "fmt"

func ExampleRectangle() {
	r1 := Rectangle{Origin: Point{0, 0}, Extent: Point{-1, -1}}
	r2 := P(1, 2).Extent(P(3, 4))
	fmt.Println(r1)
	fmt.Println(r2)
	// Output:
	// Rect( P(0,0), P(-1,-1) )
	// Rect( P(1,2), P(3,4) )
}
