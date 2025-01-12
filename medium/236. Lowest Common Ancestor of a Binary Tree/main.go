package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

// Алгоритм
// Используем рекурсивный обход дерева.
// Если корень равен nil, p или q, возвращаем корень.
// Рекурсивно вызываем функцию для левого и правого поддеревьев.
// Если оба поддерева не равны nil, возвращаем корень.
// Если левое поддерево не равно nil, возвращаем левое поддерево.
// Иначе возвращаем правое поддерево.
// Сложность:
// Время: O(n), где n - количество узлов в дереве.
// Память: O(n), где n - количество узлов в дереве.
