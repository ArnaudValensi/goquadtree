package goquadtree

type Rect struct {
	TopLeft		Position
	// TopRight	Position
	// BottomLeft	Position
	BottomRight	Position
}

func NewRect(topLeft *Position, bottomRight *Position) *Rect {
	return &Rect {
		*topLeft,
		*bottomRight,
	}
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
