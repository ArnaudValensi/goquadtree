package goquadtree

type QuadTreeNode struct {
	parentNode	*QuadTreeNode
	topRightNode	*QuadTreeNode
	topLeftNode	*QuadTreeNode
	bottomRightNode	*QuadTreeNode
	bottomLeftNode	*QuadTreeNode

	isPartitioned	bool
	items		*List.list	//PositionItem
	maxItems	int
	rect		Rect
}

// //TODO
// func NewQuadTreeNode(parent *QuadTreeNode, rect Rect, maxItems int) *QuadTreeNode {

// }

func NewQuadTreeNode(parent *QuadTreeNode, rect Rect, maxItems int) *QuadTreeNode {
	isPartitioned := false
	items := List.Init()

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
	return this.rect
}

//TODO: is it used ?
func (this *QuadTreeNode) setRect(rect *Rect) {
	this.rect = rect
}

func (this *QuadTreeNode) Insert(item *PositionItem) {
	// If partitioned, try to find child node to add to
        if !this.insertInChild(item) {
                this.items.Add(item);
                // Check if this node needs to be partitioned
                if !this.isPartitioned && this.items.Count() > this.maxItems {
			this.partition();
                }
        }
}

func (this *QuadTreeNode) insertInChild(item *PositionItem) bool {
        if !this.isPartitioned {
		return false
	}
        if this.topLeft.ContainsRect(item.Rect) {
		this.topLeft.Insert(item)
	} else if this.topRight.ContainsRect(item.Rect) {
		this.topRight.Insert(item)
	} else if this.bottomLeft.ContainsRect(item.Rect) {
		this.bottomLeft.Insert(item)
	} else if this.bottomRight.ContainsRect(item.Rect) {
		this.bottomRight.Insert(item)
	} else {
		return false; // insert in child failed
	}
        return true;
}

func (this *QuadTreeNode) PushItemDown(e *Element) {
        if (this.insertInChild(e.Value)) {
		this.items.Remove(e)
                // this.RemoveItem(i)
                return true
        } else {
		return false
	}
}

//TODO
// func (this *QuadTreeNode) PushItemUp(int i) {
//         QuadTreePositionItem<T> m = Items[i];
//         RemoveItem(i);
//         ParentNode.Insert(m);
// }

func (this *QuadTreeNode) partition() {
	midPoint := PositionAdd(this.rect.TopLeft, Rect.BottomRight)
	midPoint.Div(2)
	
        // Create the nodes
	this.topLeftNode = NewQuadTreeNode(
		this,
		NewRect(this.rect.TopLeft, midPoint), 
		this.maxItem,
		)

	firstPos := NewPosition(midPoint.X, this.rect.TopLeft.Y)
	secondPos := NewPosition(this.rect.BottomRight.X, midPoint.Y)
	this.topRightNode = NewQuadTreeNode(
		NewRect(firstPos, secondPos), 
		this.maxItem,
		)

	firstPos = NewPosition(this.rect.TopLeft.X, midPoint.Y)
	secondPos = NewPosition(midPoint.X, this.rect.BottomRight.Y)
	this.topRightNode = NewQuadTreeNode(
		NewRect(firstPos, secondPos), 
		this.maxItem,
		)

	this.topLeftNode = NewQuadTreeNode(
		this,
		NewRect(midPoint, this.rect.BottomRight), 
		this.maxItem,
		)

        this.isPartitioned = true

	// WARNING: If we cannot insert item, the item is lost.
	//          Maybe think about that.
	for e := l.Front(); e != nil; e = e.Next() {
		this.PushItemDown(e.Value)
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

func (this *QuadTreeNode) GetAllItems(itemList *List.list) {
	itemList.PushBackList(items)

	if this.isPartitioned {
		this.topLeft.GetAllItems(itemList)
		this.topRight.GetAllItems(itemList)
		this.bottomLeft.GetAllItems(itemList)
		this.bottomRight.GetAllItems(itemList)
	}
}

//TODO
//GetItem*
