package main

func largestAltitude(gain []int) int {
	max := 0
	altitude := 0
	for _, g := range gain {
		altitude += g
		if altitude > max {
			max = altitude
		}
	}
	return max
}

// Алгоритм
// Создайте переменную maxдля хранения наибольшей высоты.
// Создайте переменную altitudeдля хранения текущей высоты.
// Пройдитесь по каждому элементу gain, используя gв качестве значения:
// Добавьте gк altitude.
// Если altitudeбольше max, обновите max.
// Верните max.
// Сложность:
// Временная сложность: O(n), где n — длина массива gain.
// Мы проходим по массиву gain один раз.
// Пространственная сложность: O(1).
