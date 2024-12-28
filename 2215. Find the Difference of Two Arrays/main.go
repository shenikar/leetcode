package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)
	for _, num := range nums1 {
		map1[num] = true
	}
	for _, num := range nums2 {
		map2[num] = true
	}
	diff1 := []int{}
	for num := range map1 {
		if !map2[num] {
			diff1 = append(diff1, num)
		}
	}
	diff2 := []int{}
	for num := range map2 {
		if !map1[num] {
			diff2 = append(diff2, num)
		}
	}
	return [][]int{diff1, diff2}

}

//Алгоритм
//Сначала создайте два пустых словаря map1 и map2.
//Затем пройдите по всем элементам nums1 и добавьте их в map1.
//Затем пройдите по всем элементам nums2 и добавьте их в map2.
//Затем создайте два пустых массива diff1 и diff2.
//Затем пройдите по всем элементам map1 и добавьте те элементы, которые отсутствуют в map2, в diff1.
//Затем пройдите по всем элементам map2 и добавьте те элементы, которые отсутствуют в map1, в diff2.
//Верните массив diff1 и diff2 в качестве результата.
//Сложность:
//Время: O(n)
//Где n - это длина массива nums1 или nums2.
//Пространство: O(n)
//Где n - это длина массива nums1 или nums2.
