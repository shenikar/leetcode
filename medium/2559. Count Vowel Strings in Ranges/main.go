package main

func vowelStrings(words []string, queries [][]int) []int {
	// Создаем массив префиксных сумм
	n := len(words)
	prefix := make([]int, n+1)

	for i, word := range words {
		if isVowel(rune(word[0])) && isVowel(rune(word[len(word)-1])) {
			prefix[i+1] = prefix[i] + 1
		} else {
			prefix[i+1] = prefix[i]
		}
	}

	// Обрабатываем запросы
	res := make([]int, len(queries))
	for i, query := range queries {
		l, r := query[0], query[1]
		res[i] = prefix[r+1] - prefix[l]
	}

	return res
}

func isVowel(char rune) bool {
	return char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u'
}

// Алгоритм
// 1. Создаем массив префиксных сумм
// 2. Проходим по массиву слов
// 3. Если первая и последняя буква слова гласные, увеличиваем префиксную сумму
// 4. Обрабатываем запросы
// 5. Возвращаем результат
// Временная сложность алгоритма O(n)
// Пространственная сложность алгоритма O(n)
