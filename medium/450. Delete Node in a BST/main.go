package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		minNode := findMin(root.Right)
		root.Val = minNode.Val
		root.Right = deleteNode(root.Right, root.Val)
	}

	return root
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}

	return node
}

// Алгоритм
// 1. Если корень пуст, вернём nil.
// 2. Если ключ меньше значения корня, рекурсивно вызовем функцию для левого поддерева.
// 3. Если ключ больше значения корня, рекурсивно вызовем функцию для правого поддерева.
// 4. Иначе, если левое поддерево пусто, вернём правое поддерево.
// 5. Иначе, если правое поддерево пусто, вернём левое поддерево.
// 6. Иначе, найдём минимальный узел в правом поддереве.
// 7. Заменим значение корня минимальным значением.
// 8. Рекурсивно вызовем функцию для правого поддерева, чтобы удалить минимальный узел.
// 9. Вернём корень.
// Временная сложность алгоритма O(H), где H - высота дерева.
// Пространственная сложность алгоритма O(H), где H - высота дерева.
