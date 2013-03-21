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
		nil,
		nil,
		nil,
		nil,
		isPartitioned,
		items,
		maxItems,
		rect
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
                if !this.isPartitioned && this.items.Count() >= this.maxItems {
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

//TODO
func (this *QuadTreeNode) PushItemDown(int i) {
	// TODO: [i]
        if (this.insertInChild(this.items[i])) {
                this.RemoveItem(i)
                return true
        } else {
		return false
	}
}

//TODO
func (this *QuadTreeNode) PushItemUp(int i) {
        QuadTreePositionItem<T> m = Items[i];
        RemoveItem(i);
        ParentNode.Insert(m);
}

//TODO
func (this *QuadTreeNode) partition() {
        // // Create the nodes
        // Vector2 MidPoint = Vector2.Divide(Vector2.Add(Rect.TopLeft, Rect.BottomRight), 2.0f);
        // TopLeftNode = new QuadTreeNode<T>(this,new FRect(Rect.TopLeft, MidPoint), MaxItems);
        // TopRightNode = new QuadTreeNode<T>(this, new FRect(new Vector2(MidPoint.X, Rect.Top), new Vector2(Rect.Right, MidPoint.Y)), MaxItems);
        // BottomLeftNode = new QuadTreeNode<T>(this, new FRect(new Vector2(Rect.Left, MidPoint.Y), new Vector2(MidPoint.X, Rect.Bottom)), MaxItems);
        // BottomRightNode = new QuadTreeNode<T>(this, new FRect(MidPoint, Rect.BottomRight), MaxItems);


	midPoint := PositionAdd(this.rect.TopLeft, Rect.BottomRight)
	midPoint.Div(2)
	
        // Create the nodes
	this.topLeftNode := NewQuadTreeNode(
		this,
		NewRect(this.rect.TopLeft, midPoint), 
		this.maxItem,
		)

	firstPos := NewPosition(midPoint.X, this.rect.TopLeft.Y)
	secondPos := NewPosition(this.rect.BottomRight.X, midPoint.Y)
	this.topRightNode := NewQuadTreeNode(
		NewRect(firstPos, secondPos), 
		this.maxItem,
		)

	firstPos := NewPosition(this.rect.TopLeft.X, midPoint.Y)
	secondPos := NewPosition(midPoint.X, this.rect.BottomRight.Y)
	this.topRightNode := NewQuadTreeNode(
		NewRect(firstPos, secondPos), 
		this.maxItem,
		)

	this.topLeftNode := NewQuadTreeNode(
		this,
		NewRect(midPoint, this.rect.BottomRight), 
		this.maxItem,
		)

        IsPartitioned = true;

	//TODO
        // Try to push items down to child nodes
        int i = 0;
        while (i < Items.Count)
            {
                if (!PushItemDown(i))
                {
                    i++;
                }
        }
}

//TODO
//GetItem*
