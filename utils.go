package vectors

import "math"

func toDegrees(radians float64) float64 {
	return radians * (180 / math.Pi)
}
func toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}
