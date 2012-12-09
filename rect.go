package goquadtree

type Rect struct {
	TopLeft		Position
	TopRight	Position
	BottomLeft	Position
	BottomRight	Position
}

type Position struct {
	X		int
	Y		int
}
