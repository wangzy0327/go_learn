package main

import (
	"fmt"
	"test/tree"
)

type MyTreeNode struct {
	node *tree.TreeNode
}

func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := MyTreeNode{myNode.node.Left}
	left.postOrder()
	right := MyTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.TreeNode

	root = tree.TreeNode{Value: 3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5, nil, nil}
	root.Right.Left = new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	root.Traverse()

	fmt.Println()

	myRoot := MyTreeNode{&root}
	myRoot.postOrder()

	//root.Right.Left.Print()
	//
	//root.Print()
	//root.SetValue(100)
	//root.Print()
}
