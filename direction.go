package vectors

type Direction rune

const (
	North Direction = 'N'
	East  Direction = 'E'
	South Direction = 'S'
	West  Direction = 'W'

	InvalidDirection Direction = '?'
)

// Opposite returns the direction opposite to the caller direction
func (d Direction) Opposite() Direction {
	switch d {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	}
	return InvalidDirection
}

func (d Direction) IsHorizontal() bool {
	return (d == East || d == West)
}
func (d Direction) IsVertical() bool {
	return (d == North || d == South)
}
