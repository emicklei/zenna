package xy

import "fmt"

func ExampleLineSegment() {
	l := LineSegment{Point{}, Point{25.0, 50.0}}
	// Point{0.0,0.0}.LineSegmentTo(Point{25.0, 50.0})
	fmt.Println(l)
	// Output:
	// {{0 0} {25 50}}
}
