package xy

const (
	axisX = iota
	axisY
)

type Axis struct {
	offset float64
	xy     int // one of AxisX, AxisY
}

func AxisY(offset float64) Axis {
	return Axis{offset, axisY}
}

func AxisX(offset float64) Axis {
	return Axis{offset, axisX}
}

func Mirrored(points []Point, axis Axis) []Point {
	mirror := []Point{}
	for _, each := range points {
		if axisY == axis.xy {
			mirror = append(mirror, Point{each.X*-1 + axis.offset, each.Y})
		} else if axisX == axis.xy {
			mirror = append(mirror, Point{each.X, each.Y*-1 + axis.offset})
		}
	}
	return mirror
}
