package main

import (

)

func prefixCounting(words []string, prefix string) int {
	count := 0

	for _, word := range words {
		if len(word) >= len(prefix) && word[:len(prefix)] == prefix {
			count++
		}
	}

	return count
}