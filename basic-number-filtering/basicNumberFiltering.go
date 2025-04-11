package numberFiltering

// helpers
type Condition func(n int) bool

var IsPrime Condition = func(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func IsOdd(num int) bool  { return num%2 == 1 }
func IsEven(num int) bool { return num%2 == 0 }

func IsMultipleOfN(n int) Condition  { return func(num int) bool { return num%n == 0 } }
func IsGreaterThanN(n int) Condition { return func(num int) bool { return num > n } }
func IsLessThanN(n int) Condition    { return func(num int) bool { return num < n } }

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
