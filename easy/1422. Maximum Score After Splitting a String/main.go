package main

func maxScore(s string) int {
	maxScore := 0
	zeroCount := 0
	oneCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			oneCount++
		}
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeroCount++
		} else {
			oneCount--
		}
		maxScore = max(maxScore, zeroCount+oneCount)
	}
	return maxScore
}

// Алгортм
// 1. Подсчитываем количество единиц в строке
// 2. Проходим по строке и считаем количество нулей и единиц
// 3. Считаем максимальный результат
// 4. Возвращаем результат
// Временная сложность алгоритма O(n)
// Пространственная сложность алгоритма O(1)
