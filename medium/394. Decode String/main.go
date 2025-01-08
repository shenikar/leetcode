package main

func decodeString(s string) string {
	stack := []string{}
	numStack := []int{}
	currentNum := 0
	currentStr := ""

	for _, c := range s {
		if c >= '0' && c <= '9' {
			// Формируем число
			currentNum = currentNum*10 + int(c-'0')
		} else if c == '[' {
			// Сохраняем текущую строку и число в стеки
			stack = append(stack, currentStr)
			numStack = append(numStack, currentNum)
			currentStr = ""
			currentNum = 0
		} else if c == ']' {
			// Повторяем строку
			repeatCount := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			currentStr = temp + repeatString(currentStr, repeatCount)
		} else {
			// Добавляем текущий символ в строку
			currentStr += string(c)
		}
	}

	return currentStr
}

// Вспомогательная функция для повторения строки
func repeatString(str string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += str
	}
	return result
}

// Алгоритм
// Используем два стека: один для хранения строк, другой для хранения чисел.
// Проходим по строке.
// Если текущий символ - число, формируем число.
// Если текущий символ - открывающая скобка, сохраняем текущую строку и число в стеки.
// Если текущий символ - закрывающая скобка, повторяем текущую строку нужное количество раз.
// Если текущий символ - буква, добавляем его в текущую строку.
// В конце возвращаем текущую строку.
// Сложность:
// Время: O(n), где n - длина строки s.
// Память: O(n), где n - длина строки s.
