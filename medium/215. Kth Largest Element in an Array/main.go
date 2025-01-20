package main

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	pivotIndex := partition(nums, left, right)

	if pivotIndex == k {
		return nums[pivotIndex]
	} else if pivotIndex > k {
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		return quickSelect(nums, pivotIndex+1, right, k)
	}
}

func partition(nums []int, left, right int) int {
	pivotValue := nums[right]
	storeIndex := left

	for i := left; i < right; i++ {
		if nums[i] < pivotValue {
			nums[i], nums[storeIndex] = nums[storeIndex], nums[i]
			storeIndex++
		}
	}

	nums[storeIndex], nums[right] = nums[right], nums[storeIndex]
	return storeIndex
}

// Объяснение алгоритма:
// Quickselect использует стратегию "разделяй и властвуй", чтобы сократить объём работы:

// Выбрать опорный элемент (pivot):

// Обычно выбирается последний элемент массива.
// Разделить массив:

// Все элементы меньше опорного перемещаются влево, а больше — вправо.
// Определить позицию опорного элемента (pivotIndex):

// После разделения опорный элемент оказывается на своей правильной позиции, как если бы массив был отсортирован.
// Рекурсивный выбор:

// Если опорный элемент оказался на позиции
// 𝑘
// k, мы нашли нужный элемент.
// Если
// 𝑘<pivotIndex
// k<pivotIndex, рекурсивно ищем в левой части массива.
// Если
// 𝑘>pivotIndex
// k>pivotIndex, рекурсивно ищем в правой части массива.
