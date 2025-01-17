package main

func minimumLength(s string) (ans int) {
	cnt := [26]int{}
	for _, c := range s {
		cnt[c-'a']++
	}
	for _, x := range cnt {
		if x > 0 {
			if x&1 == 1 {
				ans += 1
			} else {
				ans += 2
			}
		}
	}
	return
}

// Алгоритм
// Инициализируем массив cnt размером 26.
// Проходим по всем символам строки s:
// Увеличиваем значение cnt[c - 'a'] на 1.
// Проходим по всем элементам массива cnt:
// Если значение больше 0:
// Если значение нечетное, увеличиваем ans на 1.
// Иначе увеличиваем ans на 2.
// Возвращаем ans.
// Сложность:
// Время: O(n), где n - длина строки s.
// Память: O(1).
