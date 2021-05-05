package vectors

import "math"

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

func (v Vector) Resolve() (SimpleVector, SimpleVector) {
	var xDirection Direction
	if v.StartDirection.IsHorizontal() {
		xDirection = v.StartDirection
	} else if v.TowardDirection.IsHorizontal() {
		xDirection = v.TowardDirection
	}

	var yDirection Direction
	if v.StartDirection.IsVertical() {
		yDirection = v.StartDirection
	} else if v.TowardDirection.IsVertical() {
		yDirection = v.TowardDirection
	}

	svX := SimpleVector{
		Mag:       math.Sin(toRadians(v.RelativeAngle)) * v.Mag,
		Direction: xDirection,
	}
	svY := SimpleVector{
		Mag:       math.Cos(toRadians(v.RelativeAngle)) * v.Mag,
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

// AddSimple adds 2 SimpleVectors together
func (sv1 SimpleVector) Add(sv2 SimpleVector) Vector {
	// Fast track if are on the same plane
	if sv1.Direction == sv2.Direction {
		sv1.Mag += sv2.Mag
		return sv1.ToVector()
	}

	// Resultant
	R := Vector{
		// Calculate the hypotenuse (resultant vector mag)
		Mag: math.Sqrt(sv1.Mag*sv1.Mag + sv2.Mag*sv2.Mag),

		// Calculate the angle angle between vector1 and the hypotenuse
		StartDirection:  sv1.Direction,
		RelativeAngle:   toDegrees(math.Atan2(sv2.Mag, sv1.Mag)),
		TowardDirection: sv2.Direction,
	}
	return R
}

func Add(vectors []Vector, simpleVectors []SimpleVector) Vector {
	// Nil slices are of len 0
	if len(simpleVectors) == 0 {
		if len(vectors) == 0 {
			return Vector{}
		}
		simpleVectors = make([]SimpleVector, len(vectors))
	} else {
		tmp := make([]SimpleVector, len(simpleVectors)+len(vectors)*2)
		copy(tmp, simpleVectors)
		simpleVectors = tmp
	}

	for _, v := range vectors {
		dX, dY := v.Resolve()
		simpleVectors = append(simpleVectors, dX, dY)
	}

	for _, sv := range simpleVectors {
		// TODO: add all together, maybe recursively?
	}
	// TODO: return the result
}
