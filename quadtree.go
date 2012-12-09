package goquadtree

import (
	"container/list"
)

type QuadTree struct {
	rootNode	*QuadTreeNode
	worldRect	Rect
	maxItems	int
}

func NewQuadTree(worldRect Rect, maxItems int) *QuadTree {
	rootNode := NewQuadTreeNode(nil, worldRect, maxItems)
	
	return &QuadTree {
		rootNode,
		worldRect,
		maxItems
	}
}

func (this *QuadTree) GetWorldRect() Rect {
	return this.worldRect
}

func (this *QuadTree) Insert(item *PositionItem) {

}

func (this *QuadTree) Remove() {

}

// Gets a list of items containing a specified point
func (this *QuadTree) GetItemsFromPoint(point *Position, itemsList *List) {

}

// Gets a list of items intersecting a specified rectangle
func (this *QuadTree) GetItemsFromRect(rect *Rect, itemsList *List) {

}
