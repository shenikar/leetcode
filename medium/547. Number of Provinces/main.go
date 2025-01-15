package main

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	count := 0
	for i := 0; i < len(isConnected); i++ {
		if !visited[i] {
			dfs(isConnected, visited, i)
			count++
		}
	}
	return count
}

func dfs(isConnected [][]int, visited []bool, city int) {
	visited[city] = true
	for i := 0; i < len(isConnected); i++ {
		if isConnected[city][i] == 1 && !visited[i] {
			dfs(isConnected, visited, i)
		}
	}
}

// Алгоритм
// 1. Инициализируем массив visited, чтобы отслеживать посещённые города.
// 2. Проходим по всем городам.
// 3. Если город не посещён, вызываем функцию dfs для посещения всех городов.
// 4. Увеличиваем счётчик групп.
// 5. Возвращаем количество групп.
// Временная сложность алгоритма O(N^2), где N - количество городов.
// Пространственная сложность алгоритма O(N), где N - количество городов.
