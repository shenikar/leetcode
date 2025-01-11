package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// Шаг 1: Найти середину списка
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Шаг 2: Развернуть вторую половину списка
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// Шаг 3: Найти максимальную сумму близнецов
	maxSum := 0
	first, second := head, prev
	for second != nil {
		sum := first.Val + second.Val
		if sum > maxSum {
			maxSum = sum
		}
		first = first.Next
		second = second.Next
	}

	return maxSum
}

// Алгоритм
// Используем два указателя: медленный и быстрый.
// Проходим по списку, пока быстрый указатель не достигнет конца списка.
// При каждом шаге медленный указатель сдвигается на один элемент, а быстрый на два.
// Когда быстрый указатель достигнет конца списка, медленный указатель будет указывать на середину списка.
// Развернем вторую половину списка.
// Проходим по списку и находим максимальную сумму близнецов.
// Возвращаем максимальную сумму близнецов.
// Сложность:
// Время: O(n), где n - количество узлов в списке.
// Память: O(1).
