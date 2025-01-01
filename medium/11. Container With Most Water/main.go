package main

func maxArea(height []int) int {
	maxArea := 0
	left, right := 0, len(height)-1
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// Алгоритм
// Инициализируйте maxAreaкак 0.
// Инициализируйте leftи rightкак 0 и длину массива height - 1.
// Пока left < right:
// Вычислите площадь между leftи right.
// Обновите maxArea.
// Если height[left] < height[right]:
// Увеличьте left.
// Иначе:
// Уменьшите right.
// Верните maxArea.
// Сложность:
// Время: O(n)
// Где n - это длина массива height.
// Мы проходим по массиву height только один раз.
// Пространство: O(1)
// Мы используем константное количество дополнительной памяти.
