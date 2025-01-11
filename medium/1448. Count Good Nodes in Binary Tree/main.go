package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, max int) int {
	if node == nil {
		return 0
	}

	good := 0
	if node.Val >= max {
		good++
		max = node.Val
	}

	return good + dfs(node.Left, max) + dfs(node.Right, max)
}

// Алгоритм
// Используем глубокий поиск в дереве.
// Передаем максимальное значение в поддеревья.
// Если значение узла больше или равно максимальному значению, увеличиваем счетчик хороших узлов.
// Вызываем рекурсивно функцию для левого и правого поддеревьев.
// Возвращаем сумму хороших узлов.
// Сложность:
// Время: O(n), где n - количество узлов в дереве.
// Память: O(h), где h - высота дерева.
