package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// Алгоритм
// Инициализируйте prevпеременную с nil.
// Пройдитесь по списку, пока headне станет nil:
// Сохраните следующий узел в next.
// Установите следующий узел текущего узла в prev.
// Переместите prevна текущий узел.
// Переместите текущий узел на следующий узел.
// Верните prevв качестве результата.
// Сложность:
// Временная сложность: O(n), где n — количество узлов в списке.
// Мы проходим по всем узлам списка.
// Пространственная сложность: O(1).
// Мы используем постоянное количество дополнительной памяти.
