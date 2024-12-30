package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := []int{}
	leaves2 := []int{}
	dfs(root1, &leaves1)
	dfs(root2, &leaves2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i := 0; i < len(leaves1); i++ {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}
	return true
}

func dfs(node *TreeNode, leaves *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*leaves = append(*leaves, node.Val)
	}
	dfs(node.Left, leaves)
	dfs(node.Right, leaves)
}

// Алгоритм
// Инициализируйте два пустых массива leaves1и leaves2для хранения листьев двух деревьев.
// Вызовите функцию dfsдля обоих деревьев, чтобы найти все листья.
// Проверьте, что количество листьев в обоих деревьях одинаково.
// Проверьте, что все листья в обоих деревьях одинаковы.
// Верните true, если все условия выполнены, иначе верните false.
// Сложность:
// Временная сложность: O(n + m), где n и m — количество узлов в двух деревьях.
// Мы проходим по всем узлам обоих деревьев.
// Пространственная сложность: O(n + m).
// Мы используем два массива для хранения листьев двух деревьев.
// func dfs это вспомогательная функция для поиска всех листьев в дереве.
// Входные параметры:
// node: текущий узел.
// leaves: массив для хранения листьев.
