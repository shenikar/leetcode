package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[i]

			if i == size-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
	}

	return result
}

// Алгоритм
// Используем обход в ширину.
// Инициализируем массив result.
// Инициализируем очередь queue и добавляем в нее корень.
// Пока очередь не пуста:
// Инициализируем переменную size равную длине очереди.
// Проходим по всем элементам очереди:
// Если i равно size-1, добавляем значение узла в result.
// Если у узла есть левый потомок, добавляем его в очередь.
// Если у узла есть правый потомок, добавляем его в очередь.
// Обрезаем очередь от 0 до size.
// Возвращаем result.
// Сложность:
// Время: O(n), где n - количество узлов в дереве.
// Память: O(n), где n - количество узлов в дереве.
