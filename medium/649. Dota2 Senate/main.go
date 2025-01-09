package main

func predictPartyVictory(senate string) string {
	radiant := []int{}
	dire := []int{}

	for i, c := range senate {
		if c == 'R' {
			radiant = append(radiant, i)
		} else {
			dire = append(dire, i)
		}
	}

	for len(radiant) > 0 && len(dire) > 0 {
		if radiant[0] < dire[0] {
			radiant = append(radiant, radiant[0]+len(senate))
		} else {
			dire = append(dire, dire[0]+len(senate))
		}
		radiant = radiant[1:]
		dire = dire[1:]
	}

	if len(radiant) > 0 {
		return "Radiant"
	}

	return "Dire"
}

// Алгоритм
// Создаем два стека: один для хранения индексов радиантов, другой для хранения индексов дир.
// Проходим по строке и добавляем индексы в соответствующие стеки.
// Пока в обоих стеках есть элементы:
// Если индекс радиантов меньше индекса дир, то добавляем индекс радиантов в конец стека радиантов.
// Иначе добавляем индекс дир в конец стека дир.
// Удаляем первый элемент из обоих стеков.
// Возвращаем результат.
// Сложность:
// Время: O(n), где n - длина строки senate.
// Память: O(n), где n - длина строки senate.
