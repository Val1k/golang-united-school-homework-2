package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesCircle   intCustom = 0
	SidesTriangle intCustom = 3
	SidesSquare   intCustom = 4
)

type intCustom int

func CalcSquare(sideLen float64, sidesNum intCustom) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (math.Sqrt(3) * math.Pow(sideLen, 2)) / 4
	case SidesSquare:
		return math.Pow(2, 2)
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)
	default:
		return 0
	}
}
