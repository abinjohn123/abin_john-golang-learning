package main

import "fmt"

func main() {
	// initialise variables
	// var smsSendingLimit int
	// var costPerSms float64
	// var hasPermission bool
	// var username string

	// limitShortHand := 10
	// costShortHand := 0.05
	// permissionShortHand := true
	// usernameShortHand := "john_doe"

	// fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSms, hasPermission, username)

	// fmt.Printf("%v, %f, %v, %q\n", limitShortHand, costShortHand, permissionShortHand, usernameShortHand)

	penniesPerText := 2.00

	fmt.Printf("Type of penniesPerText is %T\n", penniesPerText)
	fmt.Printf("A penny saved is %f pennies earned\n", penniesPerText)
}