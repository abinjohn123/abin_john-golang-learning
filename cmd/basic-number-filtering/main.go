package main

import (
	"fmt"

	numberFiltering "golang-learning/basic-number-filtering"
)

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	
	evenNumbers := numberFiltering.EvenNumbersFiltering(numbers)

	fmt.Printf("Hi %v", evenNumbers)
}