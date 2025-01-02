package main

import (

)

func findLHS(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	res := 0
	for num, count := range m {
		if m[num+1] > 0 {
			res = max(res, count+m[num+1])
		}
	}
	return res
}