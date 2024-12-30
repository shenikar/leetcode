package main

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

// Алгоритм
// Используйте массив dpдля хранения минимальной стоимости для каждого шага.
// Инициализируйте dpдля хранения n+1 элементов.
// Пройдитесь по числам от 2 до n:
// Обновите dp[i], чтобы хранить минимальную стоимость для i-го шага.
// Верните dp[n].
// Сложность:
// Время: O(n)
// Где n - это длина cost.
// Пространство: O(n)
// Мы используем массив dpдля хранения минимальной стоимости для каждого шага.
