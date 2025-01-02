package main

func longestSubarray(nums []int) int {
	maxLength := 0
	start, zeroCount := 0, 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[start] == 0 {
				zeroCount--
			}
			start++
		}
		maxLength = max(maxLength, end-start)
	}
	return maxLength
}

// Алгоритм
// 1. Инициализируем переменные maxLength, start, zeroCount
// 2. Проходим по массиву
// 3. Если элемент равен 0, увеличиваем zeroCount
// 4. Если zeroCount > 1, уменьшаем zeroCount
// 5. Обновляем maxLength
// 6. Возвращаем maxLength
// Временная сложность алгоритма O(n)
// Пространственная сложность алгоритма O(1)
