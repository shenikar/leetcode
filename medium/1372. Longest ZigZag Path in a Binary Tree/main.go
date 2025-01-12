package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	var maxLen int
	dfs(root, true, 0, &maxLen)
	dfs(root, false, 0, &maxLen)
	return maxLen
}

func dfs(node *TreeNode, isLeft bool, length int, maxLen *int) {
	if node == nil {
		return
	}

	// Обновляем максимальную длину
	if length > *maxLen {
		*maxLen = length
	}

	// Рекурсивно идем влево и вправо
	if isLeft {
		dfs(node.Left, true, 1, maxLen)          // Сбрасываем длину при смене направления
		dfs(node.Right, false, length+1, maxLen) // Продолжаем зигзаг вправо
	} else {
		dfs(node.Left, true, length+1, maxLen) // Продолжаем зигзаг влево
		dfs(node.Right, false, 1, maxLen)      // Сбрасываем длину при смене направления
	}
}

// Алгоритм
// Используем рекурсивный обход дерева.
// Передаем в функцию длину текущего пути и флаг, указывающий на направление.
// Если текущая длина больше максимальной, обновляем максимальную длину.
// Рекурсивно вызываем функцию для левого и правого поддеревьев.
// Если направление влево, сбрасываем длину при смене направления.
// Если направление вправо, сбрасываем длину при смене направления.
// Возвращаем максимальную длину.
// Сложность:
// Время: O(n), где n - количество узлов в дереве.
// Память: O(n), где n - количество узлов в дереве.
