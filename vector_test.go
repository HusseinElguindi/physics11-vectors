package vectors

import (
	"math"
	"testing"
)

func roundN(x float64, n float64) float64 {
	round := math.Pow(10, n)
	return math.Round(x*round) / round
}

func TestVecResolve(t *testing.T) {
	testCases := []struct {
		Vector Vector
		rX     SimpleVector
		rY     SimpleVector
	}{
		{
			// 10m [S 40deg E]
			Vector: Vector{10, South, 40, East},
			rX:     SimpleVector{6.4278760968653925, East},
			rY:     SimpleVector{7.660444431189781, South},
		},
		{
			// 20km [W 50deg N]
			Vector: Vector{20, West, 50, North},
			rY:     SimpleVector{15.32088886, West},
			rX:     SimpleVector{12.85575219, North},
		},
		{
			// 15km [W 80deg S]
			Vector: Vector{15, West, 80, South},
			rY:     SimpleVector{14.7721168, West},
			rX:     SimpleVector{2.604722665, South},
		},
		{
			Vector: Vector{330, North, 20, East},
			rX:     SimpleVector{112.87, East},
			rY:     SimpleVector{310.10, North},
		},
		{
			Vector: Vector{240, West, 11, North},
			rX:     SimpleVector{235.59, West},
			rY:     SimpleVector{45.79, North},
		},
	}

	for _, tc := range testCases {
		percision := float64(1)
		rX, rY := tc.Vector.Resolve()

		t.Logf("%+v %s\n", rX, string(rX.Direction))
		t.Logf("%+v %s\n\n", rY, string(rY.Direction))

		if roundN(rX.Mag, percision) != roundN(tc.rX.Mag, percision) || roundN(rY.Mag, percision) != roundN(tc.rY.Mag, percision) {
			t.Logf("failed: %+v\n", tc.Vector)
			t.Fail()
		}
	}
}

func TestSimpleVecAdd(t *testing.T) {
	testCases := []struct {
		sv1 SimpleVector
		sv2 SimpleVector
		R   Vector
	}{
		{
			sv1: SimpleVector{1, North},
			sv2: SimpleVector{1, North},
			R:   Vector{2, North, 0, North},
		},
		{
			sv1: SimpleVector{1, North},
			sv2: SimpleVector{2, South},
			R:   Vector{1, South, 0, South},
		},
		{
			sv1: SimpleVector{5.1, East},
			sv2: SimpleVector{14, North},
			R:   Vector{14.9, East, 69.9840404, North},
		},
		{
			sv1: SimpleVector{10, East},
			sv2: SimpleVector{20, West},
			R:   Vector{10, West, 0, West},
		},
	}

	for _, tc := range testCases {
		percision := float64(3)
		R := tc.sv1.Add(tc.sv2)

		t.Logf("%+v %s %s\n", R, string(R.StartDirection), string(R.TowardDirection))

		if roundN(R.Mag, percision) != roundN(tc.R.Mag, percision) {
			t.Fail()
		}
	}
}

func TestAdd(t *testing.T) {
	dX, dY := Vector{14.9, East, 69.9840404, North}.Resolve()

	dX1, dY1 := Vector{15.2, East, 20, North}.Resolve()
	dX2, dY2 := Vector{22.4, West, 40, North}.Resolve()
	testCases := []struct {
		simpleVectors []SimpleVector
		R             Vector
	}{
		{
			simpleVectors: []SimpleVector{
				{1, North},
				{1, North},
			},
			R: Vector{2, North, 0, North},
		},
		{
			simpleVectors: []SimpleVector{
				dX,
				dY,
			},
			R: Vector{14.9, East, 69.9840404, North},
		},
		{
			simpleVectors: []SimpleVector{dX1, dY1, dX2, dY2},
			R:             Vector{19.8070, West, 81.6508, North},
		},
		{
			simpleVectors: []SimpleVector{
				{15, East},
				{30.0267, South},
			},
			R: Vector{33.5649, East, 63.4553, South},
		},
		{
			simpleVectors: []SimpleVector{
				{100, East},
				{50, West},
			},
			R: Vector{50, East, 0, East},
		},
	}

	percision := float64(3)
	eqFunc := func(a float64, b float64) bool {
		return roundN(a, percision) == roundN(b, percision)
	}
	for _, tc := range testCases {
		R := Add(false, tc.simpleVectors...)

		t.Logf("Calculated: %+v %s %s\n", R, string(R.StartDirection), string(R.TowardDirection))
		t.Logf("Answer: %+v %s %s\n", tc.R, string(tc.R.StartDirection), string(tc.R.TowardDirection))

		if !tc.R.EqualFunc(R, eqFunc) {
			t.Fail()
		}
	}
}
