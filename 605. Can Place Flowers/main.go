package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				count++
			}
		}
	}
	return count >= n
}

// Algorithm

// Instead of finding the maximum value of count that can be obtained, as done in the last approach, we can stop the process of checking the positions for planting the flowers as soon as count becomes equal to n.
//Doing this leads to an optimization of the first approach. If count never becomes equal to n, n flowers can't be planted at the empty positions.
