package numberFiltering

import (
	"reflect"
	"testing"
)

// Story 1 - Only even numbers
type onlyEvenNumbersTest struct {
	arg []int
	expected []int
}

var onlyEvenNumbersTests = []onlyEvenNumbersTest{
	{[]int{}, []int{}},
	{[]int{1, 3, 5, 7, 9}, []int{}},
	{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{2, 4, 6, 8, 10}},
}

func TestOnlyEvenNumbers (t *testing.T) {
	for _, test := range onlyEvenNumbersTests {
		if output := EvenNumbersFiltering(test.arg); !reflect.DeepEqual(output, test.expected) {
			t.Errorf("For input %v, Output %v does not match expected %v", test.arg, output, test.expected)
		}
	}

}