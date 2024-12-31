package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}
	return result
}

// Алгоритм
// Инициализируйте массив resultдлиной nдля хранения результата.
// Инициализируйте result[0]как 1.
// Пройдитесь по всем элементам массива nums, начиная с 1:
// result[i] = result[i-1] * nums[i-1].
// Инициализируйте переменную rightкак 1.
// Пройдитесь по всем элементам массива nums в обратном порядке:
// result[i] *= right.
// right *= nums[i].
// Верните result.
// Сложность:
// Время: O(n)
// Где n - это длина массива nums.
// Пространство: O(1)
// Мы используем константное количество дополнительной памяти.
