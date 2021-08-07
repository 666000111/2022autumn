package go_demo

import (
	"testing"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}
var arr []int
func TestTree(t *testing.T) {
	root := new(TreeNode)
	root.val = 1
	root.left = new(TreeNode)
	root.right = new(TreeNode)
	root.left.val = 2
	root.right.val = 3
	dfs(2, root, false)
	for num := range arr {
		println(num)
	}
}

func dfs(target int, node *TreeNode, reslove bool) {
	if node == nil {
		return
	}
	arr = append(arr, node.val)
	if target == node.val {
		reslove = true
		return
	}
	if reslove == true {
		return
	}
	dfs(target, node.left, reslove)
	dfs(target, node.right, reslove)
	arr = arr[:len(arr)-1]
	return
}