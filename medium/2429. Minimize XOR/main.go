package main


func minimizeXor(num1 int, num2 int) int {
	countBits := func(n int) int {
		count := 0
		for n > 0 {
			count += n & 1
			n >>= 1
		}
		return count
	}

	bits := countBits(num2)
	result := 0
	for i := 31; i >= 0 && bits > 0; i-- {
		if num1&(1<<i) != 0 {
			result |= 1 << i
			bits--
		}
	}
	for i := 0; i < 32 && bits > 0; i++ {
		if result&(1<<i) == 0 {
			result |= 1 << i
			bits--
		}
	}
	return result
}

func countBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

// Алгоритм
// 1. Создадим функцию countBits для подсчёта количества установленных битов.
// 2. Посчитаем количество установленных битов в числе num2.
// 3. Инициализируем переменную result.
// 4. Пройдёмся по всем битам числа num1.
// 5. Если бит установлен в числе num1 и количество установленных битов в числе num2 больше нуля, установим бит в числе result.
// 6. Пройдёмся по всем битам числа num1.
// 7. Если бит не установлен в числе result и количество установленных битов в числе num2 больше нуля, установим бит в числе result.
// 8. Вернём result.
// Временная сложность алгоритма O(1).
// Пространственная сложность алгоритма O(1).
