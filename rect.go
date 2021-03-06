package goquadtree

import (
	"fmt"
)

type Rect struct {
	TopLeft		Position
	// TopRight	Position
	// BottomLeft	Position
	BottomRight	Position
	Width		int
	Height		int
}

func NewRect(topLeft *Position, bottomRight *Position) *Rect {
	width := bottomRight.X - topLeft.X
	height := bottomRight.Y - topLeft.Y
	return &Rect {
		*topLeft,
		*bottomRight,
		width,
		height,
	}
}

func (this *Rect) Eq(other *Rect) bool {
	return this.TopLeft.Eq(&other.TopLeft) &&
		this.BottomRight.Eq(&other.BottomRight)
}

func (this *Rect) Print() {
	fmt.Printf("(%d, %d), (%d, %d)",
		this.TopLeft.X, this.TopLeft.Y,
		this.BottomRight.X, this.BottomRight.Y,
	)
}

type Position struct {
	X		int
	Y		int
}

func NewPosition(x int, y int) *Position {
	return &Position {
		x,
		y,
	}
}

func (this *Position) Eq(other *Position) bool {
	return this.X == other.X && this.Y == other.Y
}

func (this *Position) Min(pos *Position) *Position {
	var x, y int

	if this.X > pos.X {
		x = pos.X
	} else {
		x = this.X
	}

	if this.Y > pos.Y {
		y = pos.Y
	} else {
		y = this.Y
	}

	return &Position {
		x,
		y,
	}
}

func (this *Position) Max(pos *Position) *Position {
	var x, y int

	if this.X > pos.X {
		x = this.X
	} else {
		x = pos.X
	}

	if this.Y > pos.Y {
		y = this.Y
	} else {
		y = pos.Y
	}

	return &Position {
		x,
		y,
	}
}

//TODO: change name cause sometime operation is with in, sometime with Position
func (this *Position) Mult(n int) {
	this.X = this.X * n
	this.Y = this.Y * n
}

func (this *Position) Div(n int) {
	this.X = this.X / n
	this.Y = this.Y / n
}

func PositionAdd(a *Position, b *Position) *Position {
	x := a.X + b.X
	y := a.Y + b.Y

	return &Position {
		x,
		y,
	}
}

func PositionSub(a *Position, b *Position) *Position {
	x := a.X - b.X
	y := a.Y - b.Y

	return &Position {
		x,
		y,
	}
}
