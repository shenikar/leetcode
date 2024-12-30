package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}

// Алгоритм
// Если корень равен nil или корень равен val, верните корень.
// Если val меньше корня, вызовите searchBST для левого поддерева.
// В противном случае вызовите searchBST для правого поддерева.
// Сложность:
// Временная сложность: O(h), где h — высота дерева.
// В худшем случае мы можем обойти все узлы в дереве.
// Пространственная сложность: O(h).
// В худшем случае глубина рекурсии равна высоте дерева.
