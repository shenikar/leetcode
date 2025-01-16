package main

func minReorder(n int, connections [][]int) int {
	graph := make([][]int, n)
	edges := make(map[[2]int]bool)
	for _, conn := range connections {
		from, to := conn[0], conn[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
		edges[[2]int{from, to}] = true
	}
	visited := make([]bool, n)
	count := 0
	dfs(0, graph, edges, visited, &count)
	return count
}

func dfs(city int, graph [][]int, edges map[[2]int]bool, visited []bool, count *int) {
	visited[city] = true
	for _, neighbor := range graph[city] {
		if !visited[neighbor] {
			if edges[[2]int{city, neighbor}] {
				*count++
			}
			dfs(neighbor, graph, edges, visited, count)
		}
	}
}

// Алгоритм:
// 1. Создаем граф, представляющий двунаправленные соединения между городами.
// 2. Создаем карту edges для отслеживания оригинальных направлений дорог.
// 3. Выполняем DFS, начиная с города 0:
//    - Помечаем текущий город как посещенный.
//    - Для каждого соседнего города:
//      - Если сосед не посещен:
//        - Проверяем, нужно ли менять направление дороги.
//        - Если да, увеличиваем счетчик.
//        - Рекурсивно вызываем DFS для соседа.
// 4. Возвращаем итоговое количество дорог, которые нужно развернуть.
//
// Сложность:
// Время: O(n), где n - количество городов, так как мы посещаем каждый город один раз.
// Пространство: O(n) для хранения графа и других структур данных.
