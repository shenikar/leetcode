package main

import (

)

func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right - left) / 2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// Алгоритм
// Используйте двоичный поиск для поиска первой плохой версии.
// Инициализируйте leftи rightдля представления диапазона версий.
// Пока leftменьше right:
// Найдите среднее значение mid.
// Если midявляется плохой версией, обновите right.
// Если midне является плохой версией, обновите left.
// Верните left.
// Сложность:
// Время: O(log(n))
// Где n - это число n.
// Пространство: O(1)
// Нам не нужно хранить дополнительные данные.