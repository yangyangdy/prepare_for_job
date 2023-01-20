package main

func main(){

}

type TreeNode struct {
	Left *TreeNode
	Right *TreeNode
	Val int
}
var res int
var index int

func kthSmallest(root *TreeNode, k int) int {
	index = k
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		index--
		if index == 0{
			res = root.Val
			return
		}
		dfs(root.Right)
	}
}

