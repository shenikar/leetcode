package main

func longestOnes(nums []int, k int) int {
	maxLength := 0
	start, repeat := 0, 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 1 {
			repeat++
		}
		for end-start+1-repeat > k {
			if nums[start] == 1 {
				repeat--
			}
			start++
		}
		maxLength = max(maxLength, end-start+1)
	}
	return maxLength
}

// Алгоритм
// 1. Инициализируем переменные maxLength, start, repeat
// 2. Проходим по массиву
// 3. Если элемент равен 1, увеличиваем repeat
// 4. Если end-start+1-repeat > k, уменьшаем repeat
// 5. Обновляем maxLength
// 6. Возвращаем maxLength
// Временная сложность алгоритма O(n)
// Пространственная сложность алгоритма O(1)
