package main

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	dfs(rooms, visited, 0)
	for _, v := range visited {
		if !v {
			return false
		}

	}
	return true
}

func dfs(rooms [][]int, visited []bool, room int) {
	visited[room] = true
	for _, key := range rooms[room] {
		if !visited[key] {
			dfs(rooms, visited, key)
		}
	}
}

// Алгоритм
// 1. Инициализируем массив visited, чтобы отслеживать посещённые комнаты.
// 2. Вызываем функцию dfs для посещения всех комнат.
// 3. Проверяем, что все комнаты были посещены.
// Временная сложность алгоритма O(N + E), где N - количество комнат, E - количество ключей.
// Пространственная сложность алгоритма O(N), где N - количество комнат.
