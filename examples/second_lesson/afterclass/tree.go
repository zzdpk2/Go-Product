package main

import "fmt"

type TreeNode struct {
	Value    int
	Children []*TreeNode
}

type Tree interface {
	Search(value interface{})
	Insert(value interface{})
	Traverse()
}

// BinaryTree binary tree
type BinaryTree struct {
	Root *TreeNode
}

// MultiChildTree multi children tree
type MultiChildTree struct {
	Root *TreeNode
}

func (mt *MultiChildTree) Search(value int) bool {
	return mt.searchNode(mt.Root, value)
}

func (mt *MultiChildTree) searchNode(curNode *TreeNode, value int) bool {
	if curNode == nil {
		return false
	}
	if curNode.Value == value {
		return true
	}

	for i := 0; i < len(curNode.Children); i++ {
		if mt.searchNode(curNode.Children[i], value) {
			return true
		}
	}
	return false
}

func (bt *BinaryTree) Search(value int) bool {
	return bt.searchNode(bt.Root, value)
}

func (bt *BinaryTree) searchNode(curNode *TreeNode, value int) bool {
	if curNode == nil {
		return false
	}
	if curNode.Value == value {
		return true
	}
	leftResult := bt.searchNode(bt.Root.Children[0], value)
	rightResult := bt.searchNode(bt.Root.Children[1], value)

	return leftResult || rightResult
}

func (mt *MultiChildTree) Insert(value int) {
	// Create a new node
	newNode := &TreeNode{
		Value:    value,
		Children: []*TreeNode{},
	}

	if mt.Root == nil {
		mt.Root = newNode
	} else {
		mt.insertNode(mt.Root, newNode)
	}
}

func (mt *MultiChildTree) insertNode(curNode, newNode *TreeNode) {
	if newNode.Value < curNode.Value {
		if len(curNode.Children) == 0 || curNode.Children[0] == nil {
			curNode.Children = append(curNode.Children, newNode)
		} else {
			mt.insertNode(curNode.Children[0], newNode)
		}
	} else if curNode.Value == newNode.Value {
		fmt.Println("Same value! Skip insert!")
	} else {
		// if curNode.Children[len(curNode.Children)-1] == nil {
		if len(curNode.Children) <= 1 {
			curNode.Children = append(curNode.Children, newNode)
		} else if curNode.Children[1] == nil {
			curNode.Children[1] = newNode
		} else {
			mt.insertNode(curNode.Children[len(curNode.Children)-1], newNode)
		}
	}
}

func (bt *BinaryTree) Insert(value int) {
	// Create a new node
	newNode := &TreeNode{
		Value:    value,
		Children: []*TreeNode{},
	}

	if bt.Root == nil {
		bt.Root = newNode
	} else {
		bt.insertNode(bt.Root, newNode)
	}
}

func (bt *BinaryTree) insertNode(currentNode, newNode *TreeNode) {
	if newNode.Value < currentNode.Value {
		if currentNode.Children[0] == nil {
			currentNode.Children[0] = newNode
		} else {
			bt.insertNode(currentNode.Children[0], newNode)
		}
	} else {
		if currentNode.Children[1] == nil {
			currentNode.Children[1] = newNode
		} else {
			bt.insertNode(currentNode.Children[1], newNode)
		}
	}
}

// Traverse via DFS
// func (bt *BinaryTree) Traverse() {
// 	res := make([]int, 0)
// 	bt.DFS(bt.Root, res)
// 	fmt.Printf("The result is: %v", res)
// }
//

// Traverse via BFS
func (bt *BinaryTree) Traverse() []int {
	// var res []int
	// return bt.DFS(bt.Root, res)
	// bt.BFS(bt.Root, res)
	// fmt.Printf("The result is: %v", bt.BFS(bt.Root))
	return bt.BFS(bt.Root)
}

func (bt *BinaryTree) DFS(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = append(res, root.Value)
	if len(root.Children) > 0 {
		res = bt.DFS(root.Children[0], res)
	}
	if len(root.Children) > 1 {
		res = bt.DFS(root.Children[1], res)
	}
	return res
}

func (bt *BinaryTree) BFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node.Value)

		if len(node.Children) > 0 {
			queue = append(queue, node.Children[0])
		}

		if len(node.Children) > 1 {
			queue = append(queue, node.Children[1])
		}
	}
	return res
}

func main() {
	// 创建二叉树
	root := &TreeNode{
		Value: 1,
		Children: []*TreeNode{
			{
				Value: 2,
				Children: []*TreeNode{
					{Value: 4, Children: nil},
					{Value: 5, Children: nil},
				},
			},
			{
				Value:    3,
				Children: nil,
			},
		},
	}

	// 创建 BinaryTree 结构体
	binaryTree := &BinaryTree{
		Root: root,
	}

	// 执行 BFS
	res := binaryTree.Traverse()

	// 打印 BFS 结果
	fmt.Println("BFS Result:", res)
}
