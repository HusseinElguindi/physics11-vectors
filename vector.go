package vectors

import (
	"fmt"
	"math"
)

type Vector struct {
	Mag float64

	StartDirection  Direction
	RelativeAngle   float64
	TowardDirection Direction
}

type SimpleVector struct {
	Mag       float64
	Direction Direction
}
type Resolved struct {
	Dx SimpleVector
	Dy SimpleVector
}

func (v Vector) String() string {
	return fmt.Sprintf("%.3f [%s %.3f %s]", v.Mag, v.StartDirection, v.RelativeAngle, v.TowardDirection)
}
func (sv SimpleVector) String() string {
	return fmt.Sprintf("%.3f [%s]", sv.Mag, sv.Direction)
}

func (v Vector) Equal(v2 Vector) bool {
	return v.EqualFunc(v2, func(a float64, b float64) bool {
		return a == b
	})
}
func (v Vector) EqualFunc(v2 Vector, eqFunc func(float64, float64) bool) bool {
	switch false {
	case eqFunc(v.Mag, v2.Mag):
	case eqFunc(float64(v.StartDirection), float64(v2.StartDirection)):
	case eqFunc(v.RelativeAngle, v2.RelativeAngle):
	case eqFunc(float64(v.TowardDirection), float64(v2.TowardDirection)):
	default:
		return true
	}
	return false
}

func (v Vector) IsValid() bool {
	return (v.StartDirection.IsHorizontal() || v.StartDirection.IsVertical() || v.TowardDirection.IsHorizontal() || v.TowardDirection.IsVertical())
}

func (sv SimpleVector) IsValid() bool {
	return (sv.Direction.IsHorizontal() || sv.Direction.IsVertical())
}

// TODO: Write tests to test the first conditions
func (v Vector) Resolve() (SimpleVector, SimpleVector) {
	xDirection, yDirection := v.StartDirection, v.TowardDirection
	if xDirection.IsVertical() {
		xDirection, yDirection = yDirection, xDirection
	}
	if !xDirection.IsHorizontal() {
		panic("invalid vector direction")
	}

	xMag, yMag := math.Sin, math.Cos
	if v.StartDirection == xDirection {
		xMag, yMag = yMag, xMag
	}

	svX := SimpleVector{
		Mag:       xMag(toRadians(v.RelativeAngle)) * v.Mag,
		Direction: xDirection,
	}
	svY := SimpleVector{
		Mag:       yMag(toRadians(v.RelativeAngle)) * v.Mag,
		Direction: yDirection,
	}

	return svX, svY
}

// ToVector converts the passed SimpleVector to the equivelant Vector
func (sv SimpleVector) ToVector() Vector {
	return Vector{
		Mag: sv.Mag,

		StartDirection:  sv.Direction,
		RelativeAngle:   0,
		TowardDirection: sv.Direction,
	}
}

func (v Vector) Simplify() (sv SimpleVector, ok bool) {
	if v.RelativeAngle == 0 {
		return SimpleVector{v.Mag, v.StartDirection}, true
	}
	return
}

// AddSimple adds 2 SimpleVectors together
func (sv1 SimpleVector) Add(sv2 SimpleVector) Vector {
	// Fast track if are on the same plane
	if (sv1.Direction.IsHorizontal() && sv2.Direction.IsHorizontal()) || (sv1.Direction.IsVertical() && sv2.Direction.IsVertical()) {
		if sv1.Direction != sv2.Direction {
			sv2.Mag = -sv2.Mag
		}
		sv1.Mag += sv2.Mag

		if sv1.Mag < 0 {
			sv1.Direction = sv1.Direction.Opposite()
			sv1.Mag = -sv1.Mag
		}
		return sv1.ToVector()
	}

	// Resultant
	R := Vector{
		// Calculate the hypotenuse (resultant vector mag)
		Mag: math.Sqrt(sv1.Mag*sv1.Mag + sv2.Mag*sv2.Mag),

		// Calculate the angle between vector1 and the hypotenuse
		StartDirection: sv1.Direction,
		// RelativeAngle:   toDegrees(math.Abs(math.Atan(sv2.Mag / sv1.Mag))),
		RelativeAngle:   toDegrees(math.Atan2(sv2.Mag, sv1.Mag)),
		TowardDirection: sv2.Direction,
	}
	return R
}

func Add(inverseAngle bool, simpleVectors ...SimpleVector) Vector {
	if len(simpleVectors) == 0 {
		return Vector{}
	}

	comp := Resolved{}
	for _, sv := range simpleVectors {
		var addend *SimpleVector
		if sv.Direction.IsHorizontal() {
			addend = &comp.Dx
		} else if sv.Direction.IsVertical() {
			addend = &comp.Dy
		} else {
			panic("invalid vector direction")
		}

		if addend.Direction == unsetDirection {
			addend.Direction = sv.Direction
		} else if sv.Direction != addend.Direction {
			sv.Mag = -sv.Mag
		}
		addend.Mag += sv.Mag

		if addend.Mag < 0 {
			addend.Direction = addend.Direction.Opposite()
			addend.Mag = -addend.Mag
		}
	}

	switch unsetDirection {
	case comp.Dx.Direction:
		return comp.Dy.ToVector()
	case comp.Dy.Direction:
		return comp.Dy.ToVector()
	}

	if inverseAngle {
		return comp.Dy.Add(comp.Dx)
	}
	return comp.Dx.Add(comp.Dy)
}
