package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var sideMax, sameMax int
	nums := []int{5, 4, 5, 4, 4, 5, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4}
	n1 := initTree(nums)
	tranverse(n1)
	oneSideMax(n1, &sideMax)
	fmt.Println(sideMax)
	oneSameMax(n1, &sameMax)
	fmt.Println(sameMax)
}

func insertNodeToTree(tree, node *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Val >= node.Val {
		if tree.Right == nil {
			tree.Right = node
			return
		}
		insertNodeToTree(tree.Right, node)
	} else if tree.Val < node.Val {
		if tree.Left == nil {
			tree.Left = node
			return
		}
		insertNodeToTree(tree.Left, node)
	}
	return
}

func initTree(nums []int) *TreeNode {
	//TODO 构造二叉树
	tree := TreeNode{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		node := TreeNode{Val: nums[i]}
		insertNodeToTree(&tree, &node)
	}
	return &tree
}

func tranverse(n *TreeNode) {
	if n == nil {
		return
	}
	tranverse(n.Left)
	tranverse(n.Right)
	fmt.Println(n.Val)
}

func oneSideMax(root *TreeNode, minVal *int) int {
	if root == nil {
		return 0
	}
	Left := max(0, oneSideMax(root.Left, minVal))
	Right := max(0, oneSideMax(root.Right, minVal))
	*minVal = max(*minVal, Left+Right+root.Val)
	return max(Left, Right) + root.Val
}

func oneSameMax(root *TreeNode, maxVal *int) int {
	if root == nil {
		return 0
	}
	l := oneSameMax(root.Left, maxVal)
	r := oneSameMax(root.Right, maxVal)
	var al, ar int
	if root.Left != nil && root.Val == root.Left.Val {
		al = l + 1
	}
	if root.Right != nil && root.Val == root.Right.Val {
		ar = r + 1
	}
	*maxVal = max(*maxVal, al+ar)
	return max(al, ar)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
