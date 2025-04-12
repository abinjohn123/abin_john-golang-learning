package numberFiltering

func EvenNumbersFiltering(nums []int) (res []int) {
	result := []int{}

	for _, num := range nums {
		if IsEven(num) {
			result = append(result, num)
		}
	}

	return result
}

func OddNumbersFiltering(nums []int) (res []int) {
	result := []int{}

	for _, num := range nums {
		if IsOdd(num) {
			result = append(result, num)
		}
	}

	return result
}

func PrimeNumbersFiltering(nums []int) (res []int) {
	result := []int{}

	for _, num := range nums {
		if IsPrime(num) {
			result = append(result, num)
		}
	}

	return result
}

func OddPrimeNumbersFiltering(nums []int) (res []int) {
	result := []int{}

	for _, num := range nums {
		if IsPrime(num) && IsOdd(num) {
			result = append(result, num)
		}
	}

	return result
}

func EvenAndMultiplesOfFive(nums []int) (res []int) {
	isMultipleOf5 := IsMultipleOfN(5)
	result := []int{}

	for _, num := range nums {
		if IsEven(num) && isMultipleOf5(num) {
			result = append(result, num)
		}
	}

	return result
}

func OddMultiplesOfThreeGreaterThanTen(nums []int) (res []int) {
	isMultipleOf3 := IsMultipleOfN(3)
	result := []int{}

	for _, num := range nums {
		if IsOdd(num) && isMultipleOf3(num) && num > 10 {
			result = append(result, num)
		}
	}
	return result
}

func PassesAllConditions(nums []int, conditions ...Condition) (res []int) {
	result := []int{}

	for _, num := range nums {
		passesAll := true

		for _, condition := range conditions {
			if !condition(num) {
				passesAll = false
				break
			}
		}

		if passesAll {
			result = append(result, num)
		}
	}

	return result
}

func PassesAnyCondition(nums []int, conditions ...Condition) (res []int) {
	var result []int

	for _, num := range nums {
		passesAny := false

		for _, condition := range conditions {
			if condition(num) {
				passesAny = true
				break
			}
		}

		if passesAny {
			result = append(result, num)
		}
	}

	return result
}
