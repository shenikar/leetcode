package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)
	for i, equation := range equations {
		a, b := equation[0], equation[1]
		value := values[i]

		if _, exists := graph[a]; !exists {
			graph[a] = make(map[string]float64)
		}
		if _, exists := graph[b]; !exists {
			graph[b] = make(map[string]float64)
		}

		graph[a][b] = value
		graph[b][a] = 1 / value
	}
	res := make([]float64, len(queries))
	for i, query := range queries {
		start, end := query[0], query[1]
		if _, exists := graph[start]; !exists {
			res[i] = -1
			continue
		}
		if _, exists := graph[end]; !exists {
			res[i] = -1
			continue
		}
		visited := make(map[string]bool)
		res[i] = dfs(start, end, visited, graph)
	}
	return res
}

func dfs(start, end string, visited map[string]bool, graph map[string]map[string]float64) float64 {
	if start == end {
		return 1.0
	}

	visited[start] = true
	for neighbor, value := range graph[start] {
		if !visited[neighbor] {
			result := dfs(neighbor, end, visited, graph)
			if result != -1.0 {
				return value * result
			}
		}
	}
	return -1.0
}

// Алгоритм:
// 1. Создаем граф, представляющий соединения между переменными.
// 2. Выполняем dfs для каждого запроса:
//   - Если начальная и конечная переменные совпадают, возвращаем 1.0.
//  - Помечаем текущую переменную как посещенную.
// - Для каждого соседнего узла:
//  - Если сосед не посещен:
//   - Вызываем dfs для соседа и умножаем результат на значение ребра между текущей и соседн
// переменными.
// - Если результат не -1.0, возвращаем его.
// - Если результат -1.0, продолжаем искать в других путях.
// 3. Возвращаем -1.0 для каждого неподходящего запроса.
// Временная сложность алгоритма O(V^2), где V - количество переменных.
// Пространственная сложность алгоритма O(V), где V - количество переменных.
// Внимание: В этой реализации мы не учитываем знаки уравнений, поэтому для корректного пов
