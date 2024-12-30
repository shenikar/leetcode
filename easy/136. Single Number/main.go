package main

import (

)

func singleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

// Алгоритм
// Используйте переменную resultдля хранения XOR всех чисел.
// Пройдитесь по всем числам в массиве:
// Обновите result, чтобы хранить XOR всех чисел.
// Верните result.
// Сложность:
// Время: O(n)
// Где n - это длина nums.
// Пространство: O(1)
// Мы используем только одну переменную resultдля хранения XOR всех чисел.
