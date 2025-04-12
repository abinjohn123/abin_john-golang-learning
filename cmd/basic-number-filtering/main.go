package main

import (
	"fmt"
	math "golang-learning/basic-number-filtering"
)

func getInputNumbers() []int {
	nums := []int{}

	for i := 0; i < 101; i++ {
		nums = append(nums, i)
	}

	return nums
}

func main() {
	nums := getInputNumbers()

	// Use helper functions to solve Stories 1 - 6
	story1 := math.PassesAllConditions(nums, math.IsEven)
	story2 := math.PassesAllConditions(nums, math.IsOdd)
	story3 := math.PassesAllConditions(nums, math.IsPrime)
	story4 := math.PassesAllConditions(nums, math.IsPrime, math.IsOdd)
	story5 := math.PassesAllConditions(nums, math.IsEven, math.IsGreaterThanN(5))
	story6 := math.PassesAllConditions(nums, math.IsOdd, math.IsGreaterThanN(10), math.IsMultipleOfN(3))


	fmt.Printf("Inputs: %v\n", nums)

	fmt.Printf("Story 1: %v\n\n", story1);
	fmt.Printf("Story 2: %v\n\n", story2);
	fmt.Printf("Story 3: %v\n\n", story3);
	fmt.Printf("Story 4: %v\n\n", story4);
	fmt.Printf("Story 5: %v\n\n", story5);
	fmt.Printf("Story 6: %v\n\n", story6);
}