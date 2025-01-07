package main

func backspaceCompare(s string, t string) bool {
	return removeBackspaces(s) == removeBackspaces(t)
}

func removeBackspaces(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c == '#' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}

// Алгоритм
// Используем стек для хранения символов строки.
// Проходим по строке и добавляем символ в стек, если это не звездочка.
// Если встречаем звездочку, удаляем последний символ из стека.
// В конце возвращаем строку, составленную из символов стека.
// Сложность:
// Время: O(n), где n - длина строки s.
// Память: O(n), где n - длина строки s.
