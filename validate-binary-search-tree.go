package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTMM(root, nil, nil)
}

func isValidBSTMM(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isValidBSTMM(root.Left, min, root) && isValidBSTMM(root.Right, root, max)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func main() {
	root := insertIntoBST(nil, 10)
	insertIntoBST(root, 5)
	insertIntoBST(root, 15)
	insertIntoBST(root, 6)
	insertIntoBST(root, 20)
	log.Println(isValidBST(root))
}
