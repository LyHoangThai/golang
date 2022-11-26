package main

import (
	"fmt"
	"os"
	"strconv"
)

func sum(nums []string) int {
	total := 0
	for _, value := range nums {
		value, _ := strconv.Atoi(value)
		total += value
	}
	return total
}

func main() {
	fmt.Println(sum(os.Args))
}
