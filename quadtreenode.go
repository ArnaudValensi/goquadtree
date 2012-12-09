package goquadtree

type QuadTreeNode struct {
	topRight	*QuadTree
	topLeft		*QuadTree
	bottomRight	*QuadTree
	bottomLeft	*QuadTree
}

func NewQuadTreeNode() *QuadTreeNode {
	return &QuadTreeNode {
		nil,
		nil,
		nil,
		nil
	}
}