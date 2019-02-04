package binaryTree

import "testing"

func printNode(t *testing.T, tree *BinaryIntTree) {
	if tree.left != nil {
		printNode(t, tree.left)
	}
	t.Logf("%v\n", tree.value)
	if tree.right != nil {
		printNode(t, tree.right)
	}
}

func Test_binaryTree(t *testing.T) {
	tree := NewInt(4)

	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(5)
	tree.Insert(2)
	tree.Insert(6)
	tree.Insert(3)

	findTree := tree.Search(2)
	if findTree.value != 2 {
		t.Error("[Error] Search error")
	}

	findNilTree := tree.Search(100)

	if findNilTree != nil {
		t.Error("[Error] 2. Search error")
	}

	// t.Logf("%+v\n", tree)
	printNode(t, tree)
}
