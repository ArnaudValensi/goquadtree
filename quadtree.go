package goquadtree

import (
	"container/list"
)

type QuadTree struct {
	rootNode	*QuadTreeNode
}

func NewQuadTree() *QuadTree {
	rootNode := NewQuadTreeNode()
	return &QuadTree {
		rootNode
	}
}

func (this *QuadTree) Insert() {

}

func (this *QuadTree) Remove() {

}

// Gets a list of items containing a specified point
func (this *QuadTree) GetItemsFromPoint(point *Position, itemsList *List) {

}

// Gets a list of items intersecting a specified rectangle
func (this *QuadTree) GetItemsFromRect(rect *Rect, itemsList *List) {

}
