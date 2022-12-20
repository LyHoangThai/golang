package main

import "fmt"

func peaks(array []int) []int {
	pos := []int{}
	peak := []int{}
	spread := 1

	if len(array) <= 1 {
		return array
	}

	for index := 1; index <= len(array)-2; index++ {
		if array[index] > array[index-spread] && array[index] > array[index+spread] {
			pos = append(pos, index)
			peak = append(peak, array[index])
		} else if 
	}
	return peak
}

func main() {
	data := []int{1, 2, 2, 2, 1}
	fmt.Println(peaks(data))
}

// 3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3
