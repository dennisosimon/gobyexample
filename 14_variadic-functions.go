package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, "")

	sum := 0
	for _, i := range nums {
		sum += i
	}

	fmt.Println(sum)
}

func main() {
	sum(1,2)
	sum(1,2,3)

	nums := []int {1,2,3,4}
	sum(nums...)
}


