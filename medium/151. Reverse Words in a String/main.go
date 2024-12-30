package main

import (
	"strings"
)

func reverseWords(s string) string {
	trimmed := strings.Fields(s)
	words := []byte(strings.Join(trimmed, " "))
	reverse(words)
	start := 0
	for i := 0; i <= len(words); i++ {
		if i == len(words) || words[i] == ' ' {
			reverse(words[start:i])
			start = i + 1
		}
	}
	return string(words)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Алгоритм
// Используйте функцию strings.Fieldsдля разделения строки на слова.
// Объедините слова с пробелами.
// Преобразуйте строку в массив байтов.
// Обратите массив байтов.
// Инициализируйте startдля хранения начальной позиции слова.
// Пройдитесь по всем символам в массиве байтов:
// Если символ равен пробелу или это конец строки:
// Обратите слово.
// Обновите start.
// Верните строку.
// Сложность:
// Время: O(n)
// Где n - это длина строки s.
// Пространство: O(n)
// Мы используем массив байтов для хранения слов.
