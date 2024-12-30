package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

// Алгоритм
// Если корень равен nil, верните 0.
// Вычислите глубину левого поддерева, вызвав функцию maxDepth для левого поддерева.
// Вычислите глубину правого поддерева, вызвав функцию maxDepth для правого поддерева.
// Верните максимальную глубину левого и правого поддеревьев плюс 1 в качестве результата.
