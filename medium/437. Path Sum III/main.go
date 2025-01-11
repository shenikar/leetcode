package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return dfs(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func dfs(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}

	count := 0
	if node.Val == targetSum {
		count++
	}

	count += dfs(node.Left, targetSum-node.Val)
	count += dfs(node.Right, targetSum-node.Val)

	return count
}

// Алгоритм
// Используем глубокий поиск в дереве.
// Используем префиксную сумму для хранения суммы значений узлов от корня до текущего узла.
// Для каждого узла вычисляем текущую сумму значений узлов от корня до текущего узла.
// Добавляем текущую сумму в префиксную сумму.
// Проверяем, есть ли в префиксной сумме значение, равное разности текущей суммы и целевой суммы.
// Если есть, увеличиваем счетчик путей.
// Вызываем рекурсивно функцию для левого и правого поддеревьев.
// Удаляем текущую сумму из префиксной суммы.
// Возвращаем сумму путей.
// Сложность:
// Время: O(n), где n - количество узлов в дереве.
// Память: O(h), где h - высота дерева.
