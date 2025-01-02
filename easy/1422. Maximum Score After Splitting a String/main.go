package main

import (

)

func maxScore(s string) int {
	maxScore := 0
	zeroCount := 0
	oneCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			oneCount++
		}
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeroCount++
		} else {
			oneCount--
		}
		maxScore = max(maxScore, zeroCount+oneCount)
	}
	return maxScore
}