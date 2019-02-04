package binaryTree

type BinaryIntTree struct {
	value int
	left  *BinaryIntTree
	right *BinaryIntTree
}

func NewInt(value int) *BinaryIntTree {
	return &BinaryIntTree{value: value}
}

func (tree *BinaryIntTree) Search(value int) *BinaryIntTree {
	if tree.value == value {
		return tree
	} else {
		if value < tree.value && tree.left != nil {
			return tree.left.Search(value)
		} else if value > tree.value && tree.right != nil {
			return tree.right.Search(value)
		} else {
			return nil
		}
	}
}

func (tree *BinaryIntTree) Insert(value int) {
	if value <= tree.value {
		if tree.left != nil {
			tree.left.Insert(value)
		} else {
			tree.left = NewInt(value)
		}
	} else {
		if tree.right != nil {
			tree.right.Insert(value)
		} else {
			tree.right = NewInt(value)
		}
	}
}
