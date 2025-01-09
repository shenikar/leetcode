package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head
	var prev *ListNode

	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = slow.Next

	return head
}

// Алгоритм
// Используем два указателя: медленный и быстрый.
// Проходим по списку, пока быстрый указатель не достигнет конца списка.
// При каждом шаге медленный указатель сдвигается на один элемент, а быстрый на два.
// Когда быстрый указатель достигнет конца списка, медленный указатель будет указывать на середину списка.
// Удаляем узел, на который указывает медленный указатель.
// Возвращаем голову списка.
// Сложность:
// Время: O(n), где n - количество узлов в списке.
// Память: O(1).
