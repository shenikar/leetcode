package main

func xorAllNums(nums1 []int, nums2 []int) int {
	result := 0

	// If the length of nums2 is odd, XOR all elements in nums1
	if len(nums2)%2 != 0 {
		for _, num := range nums1 {
			result ^= num
		}
	}

	// If the length of nums1 is odd, XOR all elements in nums2
	if len(nums1)%2 != 0 {
		for _, num := range nums2 {
			result ^= num
		}
	}

	return result
}

// Алгоритм:
// 1. Инициализируем переменную result как 0.
// 2. Проверяем, является ли длина nums2 нечетной:
//    - Если да, выполняем XOR всех элементов nums1 с result.
// 3. Проверяем, является ли длина nums1 нечетной:
//    - Если да, выполняем XOR всех элементов nums2 с result.
// 4. Возвращаем result.

// Обоснование:
// - Если длина nums2 четная, каждый элемент nums1 будет использован четное число раз,
//   что приведет к их взаимному уничтожению при XOR.
// - Если длина nums2 нечетная, каждый элемент nums1 будет использован нечетное число раз,
//   поэтому мы должны включить их в результат.
// - То же самое применимо к элементам nums2 в зависимости от четности длины nums1.

// Сложность:
// Время: O(n + m), где n и m - длины nums1 и nums2 соответственно.
// Пространство: O(1), используется только одна дополнительная переменная.
