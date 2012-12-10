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
        // check if the world needs resizing
        if (!this.rootNode.ContainsRect(item.GetRect())) {
		min := rootNode.Rect.TopLeft.Min(item.Rect.TopLeft)
		max := rootNode.Rect.BottomRight.Max(item.Rect.BottomRight)
		min.Mult(2)
		max.Mult(2)
		rect := NewRect(min, max)
                this.Resize(rect)
        }
        rootNode.Insert(item);
}

func (this *QuadTree) Resize(newWorld *Rect) {
        // Get all of the items in the tree
        // List<QuadTreePositionItem<T>> Components = new List<QuadTreePositionItem<T>>();
	itemList := List.Init()
	this.GetAllItems(itemList);

        // Create a new head
        this.rootNode = NewQuadTreeNode(nil, newWorld, this.maxItems)

	for e := itemList.Front(); e != nil; e = e.Next() {
		this.rootNode.Insert(e.(PositionItem))
	}
}

// Gets a list of items containing a specified point
func (this *QuadTree) GetItemsFromPoint(point *Position, itemsList *List) {
        if itemsList != null {
                this.rootNode.GetItemsFromPoint(point, itemsList);
        }
}

// Gets a list of items intersecting a specified rectangle
func (this *QuadTree) GetItemsFromRect(rect *Rect, itemsList *List) {
        if itemsList != null {
                this.rootNode.GetItemsFromRect(point, itemsList);
        }
}

func (this *QuadTree) GetAllItems(itemList *List.list) {
        if itemsList != null) {
                this.rootNode.GetAllItems(itemList);
        }
}

