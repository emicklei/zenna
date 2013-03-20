package xy

const (
	AxisX = iota
	AxisY
)

type Axis struct {
	offset float64
	xy     int // one of AxisX, AxisY
}

func NewAxisY(offset float64) Axis {
	return Axis{offset, AxisY}
}

func Mirrored(points []Point, axis Axis) []Point {
	mirror := []Point{}
	for _, each := range points {
		if AxisY == axis.xy {
			mirror = append(mirror, Point{each.X*-1 + axis.offset, each.Y})
		}
	}
	return mirror
}
