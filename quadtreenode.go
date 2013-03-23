package goquadtree

import (
	"container/list"
	"fmt"
)

type QuadTreeNode struct {
	parentNode	*QuadTreeNode
	topLeftNode	*QuadTreeNode
	topRightNode	*QuadTreeNode
	bottomLeftNode	*QuadTreeNode
	bottomRightNode	*QuadTreeNode

	isPartitioned	bool
	items		*list.List	//PositionItem
	maxItems	int
	rect		Rect
}

// NewQuadTreeNode return an initialized QuadTreeNode.
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

// Print display the object.
func (this *QuadTreeNode) Print() {
	fmt.Printf("parent:%p tl:%p tr:%p bl:%p br:%p isPart:%v nbItems:%d rect:", 
		this.parentNode, 
		this.topLeftNode,
		this.topRightNode,
		this.bottomLeftNode,
		this.bottomRightNode,
		this.isPartitioned,
		this.items.Len(),
		)
	this.rect.Print()
}

// GetRect is an accessor to rect.
func (this *QuadTreeNode) GetRect() *Rect {
	return &this.rect
}

//TODO: is it used ?
func (this *QuadTreeNode) setRect(rect *Rect) {
	this.rect = *rect
}

// Insert insert a PositionItem in the node.
func (this *QuadTreeNode) Insert(item *PositionItem, depth int) {
	// If partitioned, try to find child node to add to
        if !this.insertInChild(item, depth) {
                this.items.PushBack(item);
                // Check if this node needs to be partitioned
                if !this.isPartitioned && this.items.Len() > this.maxItems {
			this.partition(depth);
                }
        }
}

// insertInChild try to insert a PositionItem into children.
// Return true if the item can be insered, false otherwise.
func (this *QuadTreeNode) insertInChild(item *PositionItem, depth int) bool {
        if !this.isPartitioned {
		return false
	}

	child, err := this.GetNode(item.GetRect())
	if !err {
		child.Insert(item, depth + 1)
	} else {
		//this.items.PushBack(item);
		return false
	}

        return true;
}

// PushItemDown insert in a child an Element which contain a PositionItem.
// If it cannot be insered into a child, it is keep by the current node
// and moved back to the items list.
// Return true if the item is insered into a child, false otherwise.
func (this *QuadTreeNode) PushItemDown(e *list.Element, depth int) bool {
        if (this.insertInChild(e.Value.(*PositionItem), depth)) {
		this.items.Remove(e)
                // this.RemoveItem(i)
                return true
        }
	this.items.MoveToBack(e)
	return false
}

//TODO
// func (this *QuadTreeNode) PushItemUp(int i) {
//         QuadTreePositionItem<T> m = Items[i];
//         RemoveItem(i);
//         ParentNode.Insert(m);
// }

// partition split the node and allocate the subnodes.
func (this *QuadTreeNode) partition(depth int) {
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
	this.bottomLeftNode = NewQuadTreeNode(
		this,
		*NewRect(firstPos, secondPos), 
		this.maxItems,
		)

	this.bottomRightNode = NewQuadTreeNode(
		this,
		*NewRect(midPoint, &this.rect.BottomRight), 
		this.maxItems,
		)

        this.isPartitioned = true

	i := 0
	for i < this.items.Len() {
		e := this.items.Front()
		if !this.PushItemDown(e, depth) {
			i++
		}
	}
}

// GetNode determine which node the object belongs to. 
// Return node, error.
// If error is true, object cannot completely fit within a child node and 
// is part of the parent node. In this case, node is nil.
func (this *QuadTreeNode) GetNode(rect *Rect) (*QuadTreeNode, bool) {
        if this.topLeftNode.ContainsRect(rect) {
		return this.topLeftNode, false
	} else if this.topRightNode.ContainsRect(rect) {
		return this.topRightNode, false
	} else if this.bottomLeftNode.ContainsRect(rect) {
		return this.bottomLeftNode, false
	} else if this.bottomRightNode.ContainsRect(rect) {
		return this.bottomRightNode, false
	}
	return nil, true // insert in child failed
}

// ContainsRect check if the node contains parameter rect.
func (this *QuadTreeNode) ContainsRect(rect *Rect) bool {
        return (rect.TopLeft.X >= this.rect.TopLeft.X &&
                rect.TopLeft.Y >= this.rect.TopLeft.Y &&
                rect.BottomRight.X <= this.rect.BottomRight.X &&
                rect.BottomRight.Y <= this.rect.BottomRight.Y)
}

// GetAllNodeRect fill the list rectList with all rect in the tree.
func (this *QuadTreeNode) GetAllNodeRect(rectList *list.List) {
	rectList.PushBack(this.rect)

	if this.isPartitioned {
		this.topLeftNode.GetAllNodeRect(rectList)
		this.topRightNode.GetAllNodeRect(rectList)
		this.bottomLeftNode.GetAllNodeRect(rectList)
		this.bottomRightNode.GetAllNodeRect(rectList)
	}
}

// GetAllItems fill the list itemList with all items in the tree.
func (this *QuadTreeNode) GetAllItems(itemList *list.List) {
	itemList.PushBackList(this.items)

	if this.isPartitioned {
		this.topLeftNode.GetAllItems(itemList)
		this.topRightNode.GetAllItems(itemList)
		this.bottomLeftNode.GetAllItems(itemList)
		this.bottomRightNode.GetAllItems(itemList)
	}
}

// Fill itemList with all items that could collide with the given Rect.
func (this *QuadTreeNode) GetItems(itemList *list.List, rect *Rect) {
	var node *QuadTreeNode = nil
	var err bool = false
	if this.isPartitioned {
		node, err = this.GetNode(rect)
		if !err {
			node.GetItems(itemList, rect)
		} else if err {
			this.GetAllItems(itemList)
		}
	}
	itemList.PushBackList(this.items);
}
