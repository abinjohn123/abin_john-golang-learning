package numberFiltering

import (
	"reflect"
	"testing"
)

type numbersFilteringTest struct {
	arg      []int
	expected []int
}

// Story 1 - Only even numbers
var onlyEvenNumbersTests = []numbersFilteringTest{
	{[]int{}, []int{}},
	{[]int{1, 3, 5, 7, 9}, []int{}},
	{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{2, 4, 6, 8, 10}},
}

func TestOnlyEvenNumbers(t *testing.T) {
	for _, test := range onlyEvenNumbersTests {
		if output := EvenNumbersFiltering(test.arg); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("For input %v, Output %v does not match expected %v", test.arg, output, test.expected)
		}
	}

}

// Story 2 - Only odd numbers
var onlyOddNumbersTests = []numbersFilteringTest{
	{[]int{}, []int{}},
	{[]int{1, 3, 5, 7, 9}, []int{1, 3, 5, 7, 9}},
	{[]int{2, 4, 6, 8, 10}, []int{}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 3, 5, 7, 9}},
}

func TestOnlyOddNumbers(t *testing.T) {
	for _, test := range onlyOddNumbersTests {
		// given
		input := test.arg
		expected := test.expected

		// when
		output := OddNumbersFiltering(test.arg)

		//then
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("For input %v, output %v does not match expectec %v", input, output, expected)
		}
	}
}

// Story 3 - Only prime numbers
var onlyPrimeNumbersTests = []numbersFilteringTest{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}, []int{1, 2, 3, 5, 7, 11}},
}

func TestOnlyPrimNumers(t *testing.T) {
	for _, test := range onlyPrimeNumbersTests {
		// given
		input := test.arg
		expected := test.expected

		// when
		output := PrimeNumbersFiltering(input)

		// then
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("For input %v, output %v does not match expected %v", input, output, expected)
		}

	}
}

// Story 4 - Only odd prime numbers
var onlyOddPrimeNumbersTests = []numbersFilteringTest{
	{[]int{}, []int{}},
	{[]int{1}, []int{1}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, []int{1, 3, 5, 7, 11, 13}},
}

func TestOnlyOddPrimeNumbers(t *testing.T) {
	for _, test := range onlyOddPrimeNumbersTests {
		// given
		input := test.arg
		expected := test.expected

		// when
		output := OddPrimeNumbersFiltering(input)

		// then
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("For input %v, output %v did not match expected %v", input, output, expected)
		}
	}
}

// Story 5 - Only even & multiples of 5
var onlyEvenAndMultiplesOfFiveTests = []numbersFilteringTest{
	{[]int{}, []int{}},
	{[]int{1, 2, 3, 4, 5}, []int{}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, []int{10, 20}},
}

func TestOnlyEvenAndMultipesOfFive(t *testing.T) {
	for _, test := range onlyEvenAndMultiplesOfFiveTests {
		// when
		input := test.arg
		expected := test.expected

		// given
		output := EvenAndMultiplesOfFive(input)

		// then
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("For input %v, output %v does not match expected %v", input, output, expected)
		}
	}
}

// Story 6 - Odd, multiples of 3, greater than 10
var oddMultiplesOfThreeGreaterThanTenTests = []numbersFilteringTest{
	{[]int{}, []int{}},
	{[]int{1, 3, 5, 7, 9}, []int{}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, []int{15}},
}

func TestOddMultiplesOfThreeGreaterThanTen(t *testing.T) {
	for _, test := range oddMultiplesOfThreeGreaterThanTenTests {
		// when
		input := test.arg
		expected := test.expected

		// given
		output := OddMultiplesOfThreeGreaterThanTen(input)

		// then
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("For input %v, output %v, does not match expected %v", input, output, expected)
		}
	}
}

// Story 7 - Passes all conditions
type ConditionFilteringTest struct {
	nums       []int
	conditions []Condition
	expected   []int
}

var isGreaterThan3 Condition = IsGreaterThanN(3)
var isGreaterThan5 Condition = IsGreaterThanN(5)
var isLessThan10 Condition = IsLessThanN(10)

var allConditionsTests = []ConditionFilteringTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []Condition{isGreaterThan3, IsOdd}, []int{5, 7, 9}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []Condition{isGreaterThan5, IsEven}, []int{6, 8}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []Condition{isLessThan10, IsOdd}, []int{1, 3, 5, 7, 9}},
}

var anyConditionTests = []ConditionFilteringTest{
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []Condition{isGreaterThan3, IsOdd}, []int{1, 3, 4, 5, 6, 7, 8, 9}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []Condition{isGreaterThan5, IsEven}, []int{2, 4, 6, 7, 8, 9}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []Condition{isLessThan10, IsOdd}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

func TestPassesAllConditions(t *testing.T) {
	for _, test := range allConditionsTests {
		// given
		input := test.nums
		conditions := test.conditions
		expected := test.expected

		// when
		output := PassesAllConditions(input, conditions...)

		// then
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("For input %v, output %v does not match expected %v", input, output, expected)
		}
	}
}

// Story 8 - Passes any condition
func TestPassesAnyCondition(t *testing.T) {
	for _, test := range anyConditionTests {
		// when
		input := test.nums
		conditions := test.conditions
		expected := test.expected

		// given
		output := PassesAnyCondition(input, conditions...)

		// then
		if !reflect.DeepEqual(output, expected) {
			t.Errorf("For input %v, output %v, does not match expected %v", input, output, expected)
		}
	}
}
