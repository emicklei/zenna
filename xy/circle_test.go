package xy

import "fmt"

func ExampleCircle() {
	c := Circle{Point{0, 0}, 1, 0, 0}
	fmt.Println(c)
	// Output:
	// Circle_(P(0,0),1)
}
