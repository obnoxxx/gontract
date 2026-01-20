package main

// gontract example using Require and Ensure

import (
	"fmt"

	"github.com/gontract/gontract"
)

func Divide(dividend float64, divisor float64) (quotient float64) {

	gontract.Require(divisor != 0, "divisor must be non-zero")

	quotient = dividend / divisor

	gontract.Ensure(quotient*divisor == dividend, "quotient calculated correctly")

	return
}

func main() {

	dividends := [...]float64{10.0, 4.0, 1.0, 0}
	var divisor = 2.0

	for _, num := range dividends {

		quotient := Divide(num, divisor)

		fmt.Printf(" %f divided by %f  is %f\n", num, divisor, quotient)
	}
}
