package xy

import "math"

var (
	Tan60  = math.Tan(math.Pi / 3.0)
	Cos30  = math.Sqrt(3) / 2
	HCos30 = math.Sqrt(3) / 4
	TwoPi  = 2 * math.Pi
)

// R is short for degreesToRadians
func R(degrees int) float64 {
	return float64(degrees) * math.Pi / 180
}
