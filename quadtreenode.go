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

func (this *QuadTreeNode) GetRect() *Rect {
	return &this.rect
}

//TODO: is it used ?
func (this *QuadTreeNode) setRect(rect *Rect) {
	this.rect = *rect
}

func (this *QuadTreeNode) Insert(item *PositionItem, depth int) {
	fmt.Printf("Depth=%d\n", depth)

	// //TODO: Warning: work just with a maxItems = 1
	// if this.items.Len() > 0 &&
	// 	item.Eq(this.items.Front().Value.(*PositionItem)) {
	// 	//TODO: add an exception
	// 	fmt.Printf("Error: an item in the same position already exist\n")
	// 	return
	// }

	// If partitioned, try to find child node to add to
        if !this.insertInChild(item, depth) {
                this.items.PushBack(item);
		fmt.Printf("Nb items: %d\n", this.items.Len())
                // Check if this node needs to be partitioned
                if !this.isPartitioned && this.items.Len() > this.maxItems {
			this.partition(depth);
                }
        }
}

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
	
        // if this.topLeftNode.ContainsRect(item.GetRect()) {
	// 	this.topLeftNode.Insert(item, depth + 1)
	// } else if this.topRightNode.ContainsRect(item.GetRect()) {
	// 	this.topRightNode.Insert(item, depth + 1)
	// } else if this.bottomLeftNode.ContainsRect(item.GetRect()) {
	// 	this.bottomLeftNode.Insert(item, depth + 1)
	// } else if this.bottomRightNode.ContainsRect(item.GetRect()) {
	// 	this.bottomRightNode.Insert(item, depth + 1)
	// } else {
	// 	fmt.Printf("Error: cannot insert, item should be across multiple nodes\n")
	// 	return false; // insert in child failed
	// }

	// fmt.Printf("==Child==\n")
	// // fmt.Printf("%+v\n", this.topLeftNode)
	// // fmt.Printf("%+v\n", this.topRightNode)
	// // fmt.Printf("%+v\n", this.bottomLeftNode)
	// // fmt.Printf("%+v\n", this.bottomRightNode)

	// this.topLeftNode.Print()
	// fmt.Printf("\n")
	// this.topRightNode.Print()
	// fmt.Printf("\n")
	// this.bottomLeftNode.Print()
	// fmt.Printf("\n")
	// this.bottomRightNode.Print()
	// fmt.Printf("\n")

	// fmt.Printf("==/Child==\n")

        return true;
}

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

func (this *QuadTreeNode) partition(depth int) {
	midPoint := PositionAdd(&this.rect.TopLeft, &this.rect.BottomRight)
	midPoint.Div(2)
	
	// fmt.Printf("==Parition==\n")

        // Create the nodes
	this.topLeftNode = NewQuadTreeNode(
		this,
		*NewRect(&this.rect.TopLeft, midPoint), 
		this.maxItems,
		)

	// fmt.Printf("topLeftNode: \n\t%+v\n\t%+v\n", &this.rect.TopLeft, midPoint)

	firstPos := NewPosition(midPoint.X, this.rect.TopLeft.Y)
	secondPos := NewPosition(this.rect.BottomRight.X, midPoint.Y)
	this.topRightNode = NewQuadTreeNode(
		this,
		*NewRect(firstPos, secondPos), 
		this.maxItems,
		)

	// fmt.Printf("topRightNode: \n\t%+v\n\t%+v\n", firstPos, secondPos)

	firstPos = NewPosition(this.rect.TopLeft.X, midPoint.Y)
	secondPos = NewPosition(midPoint.X, this.rect.BottomRight.Y)
	this.bottomLeftNode = NewQuadTreeNode(
		this,
		*NewRect(firstPos, secondPos), 
		this.maxItems,
		)

	// fmt.Printf("bottomLeftNode: \n\t%+v\n\t%+v\n", firstPos, secondPos)

	this.bottomRightNode = NewQuadTreeNode(
		this,
		*NewRect(midPoint, &this.rect.BottomRight), 
		this.maxItems,
		)

	// fmt.Printf("bottomRightNode: \n\t%+v\n\t%+v\n", midPoint, &this.rect.BottomRight)

        this.isPartitioned = true

	// fmt.Printf("partition: Nb items before push down: %d\n", this.items.Len())

	i := 0
	for i < this.items.Len() {
		e := this.items.Front()
		if !this.PushItemDown(e, depth) {
			i++
		}
	}

	// TODO: it is not an error
	if this.items.Len() > 0 {
		fmt.Printf("ERROR")
	}

	// fmt.Printf("partition: Nb items after push down: %d\n", this.items.Len())
	// fmt.Printf("==/Parition==\n")

        // Try to push items down to child nodes
        // int i = 0;
        // while (i < Items.Count)
        // {
        //         if (!this.PushItemDown(i)) {
	// 		i++;
        //         }
        // }
}

// Determine which node the object belongs to. 
//
// Return node, error
//
// If error is true, object cannot completely fit within a child node and 
// is part of the parent node. In this case, node is nil
func (this *QuadTreeNode) GetNode(rect *Rect) (*QuadTreeNode, bool) {
	// fmt.Printf("rect:%+v\n", this.rect)
        if this.topLeftNode.ContainsRect(rect) {
		return this.topLeftNode, false
	} else if this.topRightNode.ContainsRect(rect) {
		return this.topRightNode, false
	} else if this.bottomLeftNode.ContainsRect(rect) {
		return this.bottomLeftNode, false
	} else if this.bottomRightNode.ContainsRect(rect) {
		return this.bottomRightNode, false
	}

	fmt.Printf("Error: cannot insert, item should be across multiple nodes\n")
	return nil, true // insert in child failed
}

func (this *QuadTreeNode) ContainsRect(rect *Rect) bool {
	// fmt.Printf("Here\n")
	// fmt.Printf("rect:%+v\n", this.rect)
        return (rect.TopLeft.X >= this.rect.TopLeft.X &&
                rect.TopLeft.Y >= this.rect.TopLeft.Y &&
                rect.BottomRight.X <= this.rect.BottomRight.X &&
                rect.BottomRight.Y <= this.rect.BottomRight.Y)
}

func (this *QuadTreeNode) GetAllNodeRect(rectList *list.List) {
	rectList.PushBack(this.rect)

	if this.isPartitioned {
		this.topLeftNode.GetAllNodeRect(rectList)
		this.topRightNode.GetAllNodeRect(rectList)
		this.bottomLeftNode.GetAllNodeRect(rectList)
		this.bottomRightNode.GetAllNodeRect(rectList)
	}
}

func (this *QuadTreeNode) GetAllItems(itemList *list.List) {
	// _ = this.items
	itemList.PushBackList(this.items)

	if this.isPartitioned {
		this.topLeftNode.GetAllItems(itemList)
		this.topRightNode.GetAllItems(itemList)
		this.bottomLeftNode.GetAllItems(itemList)
		this.bottomRightNode.GetAllItems(itemList)
	}
}

// Fill itemList with all items that could collide with the given Rect
func (this *QuadTreeNode) GetItems(itemList *list.List, rect *Rect) {
	// fmt.Printf("rect:%+v\n", rect)

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