package main

func wordSubsets(words1 []string, words2 []string) []string {
	count := make([]int, 26)
	for _, word := range words2 {
		wordCount := countLetters(word)
		for i := 0; i < 26; i++ {
			count[i] = max(count[i], wordCount[i])
		}
	}

	result := make([]string, 0)
	for _, word := range words1 {
		wordCount := countLetters(word)
		if isSubset(wordCount, count) {
			result = append(result, word)
		}
	}

	return result
}

func countLetters(word string) []int {
	count := make([]int, 26)
	for _, ch := range word {
		count[ch-'a']++
	}
	return count
}

func isSubset(wordCount []int, count []int) bool {
	for i := 0; i < 26; i++ {
		if wordCount[i] < count[i] {
			return false
		}
	}
	return true
}

// Алгоритм
// Создаем массив count, в котором будем хранить максимальное количество каждой буквы из words2.
// Проходим по всем словам из words2 и для каждого слова считаем количество каждой буквы.
// Обновляем массив count, хранящий максимальное количество каждой буквы.
// Создаем массив result, в котором будем хранить слова из words1, являющиеся подмножеством words2.
// Проходим по всем словам из words1 и для каждого слова считаем количество каждой буквы.
// Проверяем, является ли слово подмножеством words2.
// Если является, добавляем слово в массив result.
// Возвращаем массив result.
// Сложность:
// Время: O(n + m), где n - количество слов в words1, m - количество слов в words2.
// Память: O(1).
