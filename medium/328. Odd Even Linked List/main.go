package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd, even := head, head.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead

	return head
}

// Алгоритм
// Используем два указателя: один для нечетных узлов, другой для четных.
// Сохраняем начало четных узлов.
// Проходим по списку, пока четный указатель и следующий за ним узел не равны nil.
// Перемещаем нечетный указатель на следующий нечетный узел.
// Перемещаем четный указатель на следующий четный узел.
// Перемещаем указатель на следующий четный узел на следующий нечетный узел.
// Возвращаем голову списка.
// Сложность:
// Время: O(n), где n - количество узлов в списке.
// Память: O(1).
