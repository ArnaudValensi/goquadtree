// TODO Description
package goquadtree

import (
	"container/list"
)

type QuadTree struct {
	rootNode	*quadTreeNode
	worldRect	Rect
	maxItems	int
}

// newQuadTree return an initialized QuadTree.
func NewQuadTree(worldRect Rect, maxItems int) *QuadTree {
	rootNode := newQuadTreeNode(nil, worldRect, maxItems)

	return &QuadTree {
	rootNode,
	worldRect,
	maxItems,
}
}

// GetWorldRect return the world rectangle.
func (this *QuadTree) GetWorldRect() Rect {
	return this.worldRect
}

// Insert insert a PositionItem in the tree.
func (this *QuadTree) Insert(item *PositionItem) {
	// check if the world needs resizing
	itemRect := item.GetRect()
	if !this.rootNode.ContainsRect(itemRect) {
		rootNodeRect := this.rootNode.GetRect()
		min := rootNodeRect.TopLeft.Min(&itemRect.TopLeft)
		max := rootNodeRect.BottomRight.Max(&itemRect.BottomRight)
		min.Mult(2)
		max.Mult(2)
		rect := NewRect(min, max)
		this.Resize(rect)
	}
	this.rootNode.Insert(item, 1);
}

// Resize resize the tree.
func (this *QuadTree) Resize(newWorld *Rect) {
	// Get all of the items in the tree
	itemList := list.New()
	this.GetAllItems(itemList);

	// Create a new head
	this.rootNode = newQuadTreeNode(nil, *newWorld, this.maxItems)

	for e := itemList.Front(); e != nil; e = e.Next() {
		this.rootNode.Insert(e.Value.(*PositionItem), 1)
	}
}

// GetAllItems fill the list itemList with all items in the tree.
func (this *QuadTree) GetAllItems(itemList *list.List) {
	if itemList != nil {
		this.rootNode.GetAllItems(itemList);
	}
}

// GetAllNodeRect fill the list rectList with all rect in the tree.
func (this *QuadTree) GetAllNodeRect(rectList *list.List) {
	if rectList != nil {
		this.rootNode.GetAllNodeRect(rectList)
	}
}

// Fill itemList with all items that could collide with the given Rect.
func (this *QuadTree) GetItems(itemList *list.List, rect *Rect) {
	if itemList != nil && rect != nil {
		this.rootNode.GetItems(itemList, rect)
	}
}
