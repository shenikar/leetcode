package main
// Обзор
// Нам дан целочисленный массив candies, где каждое число candies[i]представляет количество конфет.я 
// т ч
//  у ребенка есть, и целое число extraCandies, обозначающее количество дополнительных конфет, которые у вас есть.

// Наша задача — вернуть логический массив resultдлины n, где result[i]true, если после указанияя 
// т ч
//  ребенок всех extraCandies, у них будет наибольшее количество конфет среди всех детей, или falseнет.

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, c := range candies {
		if c > maxCandies {
			maxCandies = c
		}
	}
	res := make([]bool, len(candies))
	for i, c := range candies {
		if c+extraCandies >= maxCandies {
			res[i] = true
		}

	}
	return res
}


// Алгоритм
// Создайте целочисленную переменную с именем maxCandiesдля хранения наибольшего количества конфет в candies. Инициализируем ее с помощью 0.
// Мы повторяем итерацию candies, и для каждого ребенка, у которого есть candyконфеты, мы выполняем действие maxCandies = max(maxCandies, candy), чтобы получить наибольшее количество конфет за candies.
// Создайте логический список answer.
// Мы повторяем этот процесс candiesеще раз и для каждого ребенка, у которого есть candyконфеты, мы прибавляем candy + extraCandies >= maxCandiesк answer.
// Возвращаться answer.