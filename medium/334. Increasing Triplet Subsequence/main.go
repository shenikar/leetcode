package main

func increasingTriplet(nums []int) bool {
	first, second := int(^uint(0)>>1), int(^uint(0)>>1)
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}

// Алгоритм
// Инициализируйте firstи secondкак максимальные значения int.
// Пройдитесь по всем элементам массива nums:
// Если numменьше или равен first:
// Обновите first.
// В противном случае, если numменьше или равен second:
// Обновите second.
// В противном случае:
// Верните true.
// Верните false.
// Сложность:
// Время: O(n)
// Где n - это длина массива nums.
// Пространство: O(1)
// Мы используем константное количество дополнительной памяти.
