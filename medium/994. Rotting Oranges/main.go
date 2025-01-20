package main

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	queue := [][2]int{}
	fresh := 0

	// Находим все гнилые апельсины и считаем свежие
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			} else if grid[i][j] == 1 {
				fresh++
			}
		}
	}

	// Направления для соседних клеток
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	minutes := 0

	// BFS
	for len(queue) > 0 && fresh > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cell := queue[0]
			queue = queue[1:]

			for _, dir := range directions {
				row, col := cell[0]+dir[0], cell[1]+dir[1]
				if row >= 0 && row < m && col >= 0 && col < n && grid[row][col] == 1 {
					grid[row][col] = 2
					fresh--
					queue = append(queue, [2]int{row, col})
				}
			}
		}
		minutes++
	}

	if fresh > 0 {
		return -1
	}
	return minutes
}

// Объяснение алгоритма:
// 1.
// Сначала мы проходим по всей сетке, чтобы найти все гнилые апельсины (2) и посчитать количество свежих апельсинов (1). Гнилые апельсины добавляются в очередь для BFS.
// 2.
// Мы используем BFS для распространения гниения. На каждом шаге:
// Обрабатываем все гнилые апельсины в текущем "уровне" (минуте).
// Для каждого гнилого апельсина проверяем его соседей в четырех направлениях.
// Если сосед - свежий апельсин, он становится гнилым, добавляется в очередь, и счетчик свежих апельсинов уменьшается.
// 3.
// После каждого "уровня" BFS увеличиваем счетчик минут.
// 4.
// Процесс продолжается, пока есть гнилые апельсины в очереди и остаются свежие апельсины.
// 5.
// В конце:
// Если остались свежие апельсины, возвращаем -1 (невозможно сгноить все апельсины).
// Иначе возвращаем количество прошедших минут.
// Временная сложность: O(m * n), где m и n - размеры сетки.
// Пространственная сложность: O(m * n) в худшем случае, когда все апельсины гнилые.