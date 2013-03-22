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

	//TODO: Warning: work just with a maxItems = 1
	if this.items.Len() > 0 &&
		item.Eq(this.items.Front().Value.(*PositionItem)) {
		//TODO: add an exception
		fmt.Printf("Error: an item in the same position already exist\n")
		return
	}

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

        if this.topLeftNode.ContainsRect(item.GetRect()) {
		this.topLeftNode.Insert(item, depth + 1)
	} else if this.topRightNode.ContainsRect(item.GetRect()) {
		this.topRightNode.Insert(item, depth + 1)
	} else if this.bottomLeftNode.ContainsRect(item.GetRect()) {
		this.bottomLeftNode.Insert(item, depth + 1)
	} else if this.bottomRightNode.ContainsRect(item.GetRect()) {
		this.bottomRightNode.Insert(item, depth + 1)
	} else {
		return false; // insert in child failed
	}

	fmt.Printf("==Child==\n")
	// fmt.Printf("%+v\n", this.topLeftNode)
	// fmt.Printf("%+v\n", this.topRightNode)
	// fmt.Printf("%+v\n", this.bottomLeftNode)
	// fmt.Printf("%+v\n", this.bottomRightNode)

	this.topLeftNode.Print()
	fmt.Printf("\n")
	this.topRightNode.Print()
	fmt.Printf("\n")
	this.bottomLeftNode.Print()
	fmt.Printf("\n")
	this.bottomRightNode.Print()
	fmt.Printf("\n")

	fmt.Printf("==/Child==\n")

        return true;
}

func (this *QuadTreeNode) PushItemDown(e *list.Element, depth int) bool {
        if (this.insertInChild(e.Value.(*PositionItem), depth)) {
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

func (this *QuadTreeNode) partition(depth int) {
	midPoint := PositionAdd(&this.rect.TopLeft, &this.rect.BottomRight)
	midPoint.Div(2)
	
	fmt.Printf("==Parition==\n")

        // Create the nodes
	this.topLeftNode = NewQuadTreeNode(
		this,
		*NewRect(&this.rect.TopLeft, midPoint), 
		this.maxItems,
		)

	fmt.Printf("topLeftNode: \n\t%+v\n\t%+v\n", &this.rect.TopLeft, midPoint)

	firstPos := NewPosition(midPoint.X, this.rect.TopLeft.Y)
	secondPos := NewPosition(this.rect.BottomRight.X, midPoint.Y)
	this.topRightNode = NewQuadTreeNode(
		this,
		*NewRect(firstPos, secondPos), 
		this.maxItems,
		)

	fmt.Printf("topRightNode: \n\t%+v\n\t%+v\n", firstPos, secondPos)

	firstPos = NewPosition(this.rect.TopLeft.X, midPoint.Y)
	secondPos = NewPosition(midPoint.X, this.rect.BottomRight.Y)
	this.bottomLeftNode = NewQuadTreeNode(
		this,
		*NewRect(firstPos, secondPos), 
		this.maxItems,
		)

	fmt.Printf("bottomLeftNode: \n\t%+v\n\t%+v\n", firstPos, secondPos)

	this.bottomRightNode = NewQuadTreeNode(
		this,
		*NewRect(midPoint, &this.rect.BottomRight), 
		this.maxItems,
		)

	fmt.Printf("bottomRightNode: \n\t%+v\n\t%+v\n", midPoint, &this.rect.BottomRight)

        this.isPartitioned = true

	fmt.Printf("partition: Nb items before push down: %d\n", this.items.Len())

	i := 0
	for this.items.Len() > 0 {
		e := this.items.Front()
		if !this.PushItemDown(e, depth) {
			i++
		}
	}
	// TODO: if this.items.Len() => error
	if this.items.Len() > 0 {
		fmt.Printf("ERROR")
	}

	fmt.Printf("partition: Nb items after push down: %d\n", this.items.Len())
	fmt.Printf("==/Parition==\n")

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
	// _ = this.items
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
	// _ = rect.TopLeft.X >= this.rect.TopLeft.X
        // _ = rect.TopLeft.Y >= this.rect.TopLeft.Y
        // _ = rect.BottomRight.X <= this.rect.BottomRight.X
        // _ = rect.BottomRight.Y <= this.rect.BottomRight.Y

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
