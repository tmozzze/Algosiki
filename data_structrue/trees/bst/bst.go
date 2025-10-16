package bst

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func NewBST() *BST {
	return &BST{Root: nil}
}

func (t *BST) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return &Node{Key: key}
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	}

	return node
}

func (t *BST) Delete(key int) bool {
	t.Root = delete(t.Root, key)
}
