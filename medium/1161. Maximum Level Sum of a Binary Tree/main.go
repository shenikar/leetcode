package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max, maxLevel := -1<<31, 0
	queue := []*TreeNode{root}

	for level := 1; len(queue) > 0; level++ {
		sum := 0
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[i]
			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if sum > max {
			max = sum
			maxLevel = level
		}

		queue = queue[size:]
	}

	return maxLevel
}

// Алгоритм
// Используем обход в ширину.
// Инициализируем переменные max и maxLevel.
// Инициализируем очередь queue и добавляем в нее корень.
// Пока очередь не пуста:
// Инициализируем переменную sum равную 0.
// Инициализируем переменную size равную длине очереди.
// Проходим по всем элементам очереди:
// Добавляем значение узла в sum.
// Если у узла есть левый потомок, добавляем его в очередь.
// Если у узла есть правый потомок, добавляем его в очередь.
// Если sum больше max:
// Обновляем max и maxLevel.
// Обрезаем очередь от 0 до size.
// Возвращаем maxLevel.
// Сложность:
// Время: O(n), где n - количество узлов в дереве.
// Память: O(n), где n - количество узлов в дереве.
