package main

import "fmt"

func totalCost(costs []int, k int, candidates int) int {
	re := 0
	multiplier := 0
	var currentLow int
	for i := 0; i <= k; i++ {
		currentLow = costs[0]
		for _, va := range costs {
			if currentLow > va {
				currentLow = va
				multiplier += 1
				fmt.Println(currentLow, va, multiplier)
			} else if currentLow == va {
				multiplier += 1
				fmt.Println(currentLow, va, multiplier)
			}
			re += (currentLow * multiplier)
			currentLow = 9999
		}
	}
	return re
}

func main() {
	costs := []int{17, 12, 10, 2, 7, 2, 11, 2, 8}

	fmt.Println(totalCost(costs, 3, 4))
}
