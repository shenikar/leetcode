package main

func pivotIndex(nums []int) int {
	var sum, leftSum int
	for _, num := range nums {
		sum += num
	}
	for i, num := range nums {
		if leftSum == sum-leftSum-num {
			return i
		}
		leftSum += num
	}
	return -1
}

// Алгоритм
// Инициализируйте переменные sumи leftSumнулевыми значениями для хранения суммы всех элементов и суммы элементов слева от текущего индекса.
// Пройдитесь по массиву numsдля вычисления суммы всех элементов.
// Пройдитесь по массиву nums, используя iв качестве индекса:
// Если leftSumравно sum-leftSum-nums[i], верните i.
// Добавьте nums[i]к leftSum.
// Если индекс не найден, верните -1.
// Сложность:
// Временная сложность: O(n), где n — длина массива nums.
// Мы проходим по массиву nums дважды.
// Пространственная сложность: O(1).
