package main

import (
	"fmt"
)

func compress(chars []byte) int {
	write, anchor := 0, 0
	for read := 0; read < len(chars); read++ {
		if read+1 == len(chars) || chars[read] != chars[read+1] {
			chars[write] = chars[anchor]
			write++
			if read > anchor {
				for _, c := range []byte(fmt.Sprint(read - anchor + 1)) {
					chars[write] = c
					write++
				}
			}
			anchor = read + 1
		}
	}
	return write
}

// Алгоритм
// Инициализируйте writeи anchorкак 0.
// Пройдитесь по всем символам в массиве chars:
// Если символ не равен следующему символу или это последний символ:
// Запишите символ в позицию write.
// Увеличьте write.
// Если длина повторяющихся символов больше 1:
// Запишите длину повторяющихся символов в позицию write.
// Увеличьте write.
// Обновите anchor.
// Верните write.
// Сложность:
// Время: O(n)
// Где n - это длина массива chars.
// Пространство: O(1)
// Мы используем константное количество дополнительной памяти.
