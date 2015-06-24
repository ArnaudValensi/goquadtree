package goquadtree

type PositionItem struct {
	rect	Rect
}

func NewPositionItem(rect *Rect) *PositionItem {
	return &PositionItem {*rect}
}

func (this *PositionItem) GetRect() *Rect {
	return &this.rect
}

func (this *PositionItem) Eq(other *PositionItem) bool {
	return this.rect.Eq(other.GetRect())
}

func (this *PositionItem) Add() {

}

func (this *PositionItem) Count() {

}

func (this *PositionItem) ContainsRect() {

}

func (this *PositionItem) Insert() {

}
