package vec2

import "slices"

type Vec2 struct {
	X, Y int
}

var North = Vec2{0, -1}
var East = Vec2{1, 0}
var South = Vec2{0, 1}
var West = Vec2{-1, 0}
var CardinalDirs = []Vec2{North, East, South, West}

func New(x int, y int) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) InRange(x1, y1, x2, y2 int) bool {
	return v.X >= x1 && v.X <= x2 && v.Y >= y1 && v.Y <= y2
}

func (va Vec2) Add(vb Vec2) Vec2 {
	return Vec2{va.X + vb.X, va.Y + vb.Y}
}

func (va Vec2) Sub(vb Vec2) Vec2 {
	return Vec2{va.X - vb.X, va.Y - vb.Y}
}

func (v Vec2) Mul(s int) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

func (v Vec2) Div(s int) Vec2 {
	return Vec2{v.X / s, v.Y / s}
}

func (v Vec2) RotateCardinalCW() Vec2 {
	i := slices.Index(CardinalDirs, v)
	if i < 0 || i == 3 {
		return North
	}

	return CardinalDirs[i+1]
}

func (v Vec2) RotateCardinalCCW() Vec2 {
	i := slices.Index(CardinalDirs, v)
	if i < 1 {
		return West
	}

	return CardinalDirs[i-1]
}

func FromRune(c rune) Vec2 {
	switch c {
	case 'N':
		return North
	case 'U':
		return North
	case 'E':
		return East
	case 'R':
		return East
	case 'S':
		return South
	case 'D':
		return South
	case 'W':
		return West
	case 'L':
		return West
	default:
		panic("Not a valid direction")
	}
}
