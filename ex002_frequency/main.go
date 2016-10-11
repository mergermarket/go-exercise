package main

import "fmt"

func main() {
	nums := []int{10, 5, 2, 10, 3, 5, 7, 1, 2, 4, 10, 9, 9, 1, 4, 6, 1, 7, 10, 9, 6, 5, 1, 4, 10, 1, 7, 7, 5, 6, 7, 6, 7, 9, 2, 6, 4, 5, 1, 8, 8, 3, 3, 7, 2, 4, 1, 5, 9, 1, 1, 5, 9, 3, 6, 7, 1, 4, 1, 3, 2, 2, 7, 4, 7, 7, 7, 7, 10, 2, 2, 10, 6, 6, 5, 9, 1, 10, 3, 4, 8, 8, 7, 9, 8, 3, 10, 4, 8, 9, 1, 9, 6, 9, 2, 6, 1, 8, 2, 7}

	frequency, uniques := analyseFrequency(nums)

	fmt.Println(frequency) // expected: map[1:14 4:9 9:11 5:8 2:10 3:7 7:15 6:10 8:7 10:9]
	fmt.Println(uniques)   // expected: [10 5 2 3 7 1 4 9 6 8]  --- in the order in which they occur in nums
}
