package main

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	for _, num := range arr {
		count[num]++
	}
	occ := make(map[int]bool)
	for _, c := range count {
		if occ[c] {
			return false
		}
		occ[c] = true
	}
	return true
}

//Алгоритм
//Сначала создайте пустой словарь count.
//Затем пройдите по всем элементам массива arr и добавьте их в count.
//Затем создайте пустой словарь occ.
//Затем пройдите по всем значениям count и добавьте их в occ.
//Если значение уже присутствует в occ, верните false.
//Верните true в противном случае.
//Сложность:
//Время: O(n)
//Где n - это длина массива arr.
//Пространство: O(n)
//Где n - это длина массива arr.
