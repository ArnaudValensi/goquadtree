package goquadtree

import (
	"container/list"
)

type QuadTreeNode struct {
	parentNode	*QuadTreeNode
	topRightNode	*QuadTreeNode
	topLeftNode	*QuadTreeNode
	bottomRightNode	*QuadTreeNode
	bottomLeftNode	*QuadTreeNode

	isPartitioned	bool
	items		*list.List	//PositionItem
	maxItems	int
	rect		Rect
}

// //TODO
// func NewQuadTreeNode(parent *QuadTreeNode, rect Rect, maxItems int) *QuadTreeNode {

// }

func NewQuadTreeNode(parent *QuadTreeNode, rect Rect, maxItems int) *QuadTreeNode {
	isPartitioned := false
	items := list.New()

	return &QuadTreeNode {
		parent,
		nil,
		nil,
		nil,
		nil,
		isPartitioned,
		items,
		maxItems,
		rect,
	}
}

func (this *QuadTreeNode) GetRect() *Rect {
	return &this.rect
}

//TODO: is it used ?
func (this *QuadTreeNode) setRect(rect *Rect) {
	this.rect = *rect
}

func (this *QuadTreeNode) Insert(item *PositionItem) {
	// If partitioned, try to find child node to add to
        if !this.insertInChild(item) {
                this.items.PushBack(item);
                // Check if this node needs to be partitioned
                if !this.isPartitioned && this.items.Len() > this.maxItems {
			this.partition();
                }
        }
}

func (this *QuadTreeNode) insertInChild(item *PositionItem) bool {
        if !this.isPartitioned {
		return false
	}
        if this.topLeftNode.ContainsRect(item.GetRect()) {
		this.topLeftNode.Insert(item)
	} else if this.topRightNode.ContainsRect(item.GetRect()) {
		this.topRightNode.Insert(item)
	} else if this.bottomLeftNode.ContainsRect(item.GetRect()) {
		this.bottomLeftNode.Insert(item)
	} else if this.bottomRightNode.ContainsRect(item.GetRect()) {
		this.bottomRightNode.Insert(item)
	} else {
		return false; // insert in child failed
	}
        return true;
}

func (this *QuadTreeNode) PushItemDown(e *list.Element) bool {
        if (this.insertInChild(e.Value.(*PositionItem))) {
		this.items.Remove(e)
                // this.RemoveItem(i)
                return true
        }
	return false
}

//TODO
// func (this *QuadTreeNode) PushItemUp(int i) {
//         QuadTreePositionItem<T> m = Items[i];
//         RemoveItem(i);
//         ParentNode.Insert(m);
// }

func (this *QuadTreeNode) partition() {
	midPoint := PositionAdd(&this.rect.TopLeft, &this.rect.BottomRight)
	midPoint.Div(2)
	
        // Create the nodes
	this.topLeftNode = NewQuadTreeNode(
		this,
		*NewRect(&this.rect.TopLeft, midPoint), 
		this.maxItems,
		)

	firstPos := NewPosition(midPoint.X, this.rect.TopLeft.Y)
	secondPos := NewPosition(this.rect.BottomRight.X, midPoint.Y)
	this.topRightNode = NewQuadTreeNode(
		this,
		*NewRect(firstPos, secondPos), 
		this.maxItems,
		)

	firstPos = NewPosition(this.rect.TopLeft.X, midPoint.Y)
	secondPos = NewPosition(midPoint.X, this.rect.BottomRight.Y)
	this.topRightNode = NewQuadTreeNode(
		this,
		*NewRect(firstPos, secondPos), 
		this.maxItems,
		)

	this.topLeftNode = NewQuadTreeNode(
		this,
		*NewRect(midPoint, &this.rect.BottomRight), 
		this.maxItems,
		)

        this.isPartitioned = true

	// WARNING: If we cannot insert item, the item is lost.
	//          Maybe think about that.
	for e := this.items.Front(); e != nil; e = e.Next() {
		this.PushItemDown(e)
	}

        // Try to push items down to child nodes
        // int i = 0;
        // while (i < Items.Count)
        // {
        //         if (!this.PushItemDown(i)) {
	// 		i++;
        //         }
        // }
}

func (this *QuadTreeNode) GetAllItems(itemList *list.List) {
	itemList.PushBackList(this.items)

	if this.isPartitioned {
		this.topLeftNode.GetAllItems(itemList)
		this.topRightNode.GetAllItems(itemList)
		this.bottomLeftNode.GetAllItems(itemList)
		this.bottomRightNode.GetAllItems(itemList)
	}
}

//TODO
//GetItem*

func (this *QuadTreeNode) ContainsRect(rect *Rect) bool {
        return (rect.TopLeft.X >= this.rect.TopLeft.X &&
                rect.TopLeft.Y >= this.rect.TopLeft.Y &&
                rect.BottomRight.X <= this.rect.BottomRight.X &&
                rect.BottomRight.Y <= this.rect.BottomRight.Y)
}
