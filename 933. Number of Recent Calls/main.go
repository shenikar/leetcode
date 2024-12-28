package main

type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.pings = append(this.pings, t)
	for this.pings[0] < t-3000 {
		this.pings = this.pings[1:]
	}
	return len(this.pings)
}

// Алгоритм
// Создайте структуру RecentCounterс полем pingsдля хранения временных меток.
// Создайте конструктор, который возвращает новый экземпляр RecentCounter.
// Создайте метод Pingдля добавления временной метки в pings.
// Добавьте временную метку в pings.
// Удалите все временные метки, которые находятся вне временного интервала 3000 мс.
// Верните длину pingsв качестве результата.
// Сложность:
// Время: O(n)
// Где n - это количество вызовов метода Ping.
// Пространство: O(n)
// Где n - это количество вызовов метода Ping.
