package main

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, a := range asteroids {
		if a > 0 {
			stack = append(stack, a)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < -a {
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] < 0 {
				stack = append(stack, a)
			} else if stack[len(stack)-1] == -a {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return stack
}

// Алгоритм
// Используем стек для хранения астероидов.
// Проходим по массиву астероидов.
// Если астероид положительный, добавляем его в стек.
// Если астероид отрицательный, сравниваем его с последним астероидом в стеке.
// Если последний астероид в стеке положительный и меньше текущего отрицательного, удаляем его из стека.
// Если последний астероид в стеке отрицательный, добавляем текущий астероид в стек.
// Если последний астероид в стеке равен текущему отрицательному, удаляем его из стека.
// В конце возвращаем стек.
// Сложность:
// Время: O(n), где n - длина массива астероидов.
// Память: O(n), где n - длина массива астероидов.