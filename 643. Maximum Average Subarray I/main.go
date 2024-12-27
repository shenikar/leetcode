package main

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}

// Алгоритм
// Инициализируйте переменную sumс нулевым значением для хранения суммы первых k элементов.
// Пройдитесь по первым k элементам numsи добавьте их к sum.
// Инициализируйте переменную maxSum значением sum.
// Пройдитесь по оставшимся элементам nums начиная с k-го элемента:
// Добавьте nums[i]к sumи вычтите nums[i-k]из sum, чтобы получить сумму следующих k элементов.
// Обновите maxSum, если sum больше maxSum.
// Верните среднее значение maxSum, разделив его на k.
// Сложность:
// Временная сложность: O(n), где n — длина массива nums.
// Мы проходим по массиву nums дважды.
// Пространственная сложность: O(1).
