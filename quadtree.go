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
		maxItems,
	}
}

func (this *QuadTree) GetWorldRect() Rect {
	return this.worldRect
}

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

func (this *QuadTree) Resize(newWorld *Rect) {
        // Get all of the items in the tree
	itemList := list.New()
	this.GetAllItems(itemList);

        // Create a new head
        this.rootNode = NewQuadTreeNode(nil, *newWorld, this.maxItems)

	for e := itemList.Front(); e != nil; e = e.Next() {
		this.rootNode.Insert(e.Value.(*PositionItem), 1)
	}
}

func (this *QuadTree) GetAllItems(itemList *list.List) {
        if itemList != nil {
                this.rootNode.GetAllItems(itemList);
        }
}

func (this *QuadTree) GetAllNodeRect(rectList *list.List) {
	if rectList != nil {
		this.rootNode.GetAllNodeRect(rectList)
	}
}

func (this *QuadTree) GetItems(itemList *list.List, rect *Rect) {
        if itemList != nil && rect != nil {
		this.rootNode.GetItems(itemList, rect)
	}
}
