package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

//func Print(node TreeNode) {
//	fmt.Println(node.Value)
//}

func (node TreeNode) Print() {
	fmt.Print(node.Value, " ")
}

func (node *TreeNode) SetValue(Value int) {
	if node == nil {
		fmt.Println("Setting Value to nil" +
			"node. Ignored.")
		return
	}
	node.Value = Value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func CreateNode(Value int) *TreeNode {
	return &TreeNode{Value: Value}
}

//func main() {

//node := []TreeNode{
//	{Value: 3},
//	{},
//	{6, nil, &root},
//}
//fmt.Println(node)

//}
