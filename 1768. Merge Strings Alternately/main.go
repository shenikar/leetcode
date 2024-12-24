package main

// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
//If a string is longer than the other, append the additional letters onto the end of the merged string.
// Return the merged string.

// Этот алгоритм называется "слияние строк поочередно" (merge alternately). Он объединяет две строки, чередуя символы из каждой строки. 

// Пример работы алгоритма:
// - Входные строки: "abc" и "pqr"
// - Результат: "apbqcr"

// Алгоритм работает следующим образом:
// 1. Инициализируется пустая строка `result`.
// 2. Запускается цикл, который продолжается до тех пор, пока не достигнуты концы обеих строк.
// 3. В каждой итерации цикла добавляется символ из первой строки, если текущий индекс меньше длины первой строки.
// 4. Затем добавляется символ из второй строки, если текущий индекс меньше длины второй строки.
// 5. После завершения цикла возвращается объединенная строка `result`.

// Этот алгоритм полезен для задач, где требуется объединить две строки, сохраняя порядок символов из каждой строки.

func mergeAlternately(word1 string, word2 string) string {
	var result string
	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			result += string(word1[i])
		}
		if i < len(word2) {
			result += string(word2[i])
		}
	}
	return result

}
