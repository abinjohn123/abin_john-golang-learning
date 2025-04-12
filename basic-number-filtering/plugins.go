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