package goquadtree

import (
	"container/list"
	"fmt"
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
	
	this.rootNode.Print()
	fmt.Printf("\n")
}

func (this *QuadTree) Resize(newWorld *Rect) {
        // Get all of the items in the tree
        // List<QuadTreePositionItem<T>> Components = new List<QuadTreePositionItem<T>>();
	itemList := list.New()
	this.GetAllItems(itemList);

        // Create a new head
        this.rootNode = NewQuadTreeNode(nil, *newWorld, this.maxItems)

	for e := itemList.Front(); e != nil; e = e.Next() {
		this.rootNode.Insert(e.Value.(*PositionItem), 1)
	}
}

// // Gets a list of items containing a specified point
// func (this *QuadTree) GetItemsFromPoint(point *Position, itemsList *List) {
//         if itemsList != null {
//                 this.rootNode.GetItemsFromPoint(point, itemsList);
//         }
// }

// // Gets a list of items intersecting a specified rectangle
// func (this *QuadTree) GetItemsFromRect(rect *Rect, itemsList *List) {
//         if itemsList != null {
//                 this.rootNode.GetItemsFromRect(point, itemsList);
//         }
// }

func (this *QuadTree) GetAllItems(itemList *list.List) {
        if itemList != nil {
                this.rootNode.GetAllItems(itemList);
        }
}

func (this *QuadTree) GetAllNodeRect(rectList *list.List) {
	if this.rootNode != nil {
		this.rootNode.GetAllNodeRect(rectList)
	}
}