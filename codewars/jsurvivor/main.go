package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(helper(test, 3, 3))
}

func helper(arr []int, skip int, counter int) int {
	skip -= 1
	for len(arr) > 2 {
		if skip < len(arr) {
			arr = append(arr[:skip], arr[skip+1:]...)
			skip += counter - 1
		} else if skip >= len(arr) {
			skip = (skip % len(arr))
			arr = append(arr[:skip], arr[skip+1:]...)
			skip += counter - 1
		}
	}
	return arr[(skip+1)%2]
}
