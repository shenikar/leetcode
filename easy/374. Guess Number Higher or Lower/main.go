package main

func guessNumber(n int) int  {
	left, right := 1, n
	for left < right {
		mid := left + (right - left) / 2
		if guess(mid) == 0 {
			return mid
		} else if guess(mid) == 1 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func guess(num int) int {
	return 0
}

// Алгоритм
// Используйте двоичный поиск для поиска числа.
// Инициализируйте leftи rightдля представления диапазона чисел.
// Пока leftменьше right:
// Найдите среднее значение mid.
// Если guess(mid)равно 0, верните mid.
// Если guess(mid)равно 1, обновите left.
// Если guess(mid)равно -1, обновите right.
// Верните left.
// Сложность:
// Время: O(log(n))
// Где n - это число n.
// Пространство: O(1)
// Нам не нужно хранить дополнительные данные.
