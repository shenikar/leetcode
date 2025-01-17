package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[i-1] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

// Алгоритм работает следующим образом:
// Он использует два указателя: i и j.
// j проходит по всему массиву.
// i указывает на позицию, куда нужно записать следующий уникальный элемент.
// Когда находится новый уникальный элемент (отличный от предыдущего), он записывается в позицию i, и i увеличивается.
// В результате все уникальные элементы оказываются в начале массива, а i показывает их количество.
// Этот алгоритм эффективен, так как он работает за один проход по массиву (временная сложность O(n)) и не требует дополнительной памяти (пространственная сложность O(1)).
