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
			rX:     SimpleVector{15.32088886, West},
			rY:     SimpleVector{12.85575219, North},
		},
		{
			// 15km [W 80deg S]
			Vector: Vector{15, West, 80, South},
			rX:     SimpleVector{14.7721168, West},
			rY:     SimpleVector{2.604722665, South},
		},
	}

	for _, tc := range testCases {
		percision := float64(3)
		rX, rY := tc.Vector.Resolve()

		t.Logf("%+v %s\n", rX, string(rX.Direction))
		t.Logf("%+v %s\n\n", rY, string(rY.Direction))

		if roundN(rX.Mag, percision) != roundN(tc.rX.Mag, percision) || roundN(rY.Mag, percision) != roundN(tc.rY.Mag, percision) {
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
			sv1: SimpleVector{5.1, East},
			sv2: SimpleVector{14, North},
			R:   Vector{14.9, East, 69.9840404, North},
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
