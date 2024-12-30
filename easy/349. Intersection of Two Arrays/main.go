package main

func intersection(nums1 []int, nums2 []int) []int {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)
	for _, num := range nums1 {
		map1[num] = true
	}
	for _, num := range nums2 {
		map2[num] = true
	}
	intersection := []int{}
	for num := range map1 {
		if map2[num] {
			intersection = append(intersection, num)
		}
	}
	return intersection

}

//Алгоритм
//Сначала создайте два пустых словаря map1 и map2.
//Затем пройдите по всем элементам nums1 и добавьте их в map1.
//Затем пройдите по всем элементам nums2 и добавьте их в map2.
//Затем создайте пустой массив intersection.
//Затем пройдите по всем элементам map1 и добавьте те элементы, которые присутствуют в map2, в intersection.
//Верните массив intersection в качестве результата.
//Сложность:
//Время: O(n)
//Где n - это длина массива nums1 или nums2.
//Пространство: O(n)
//Где n - это длина массива nums1 или nums2.
