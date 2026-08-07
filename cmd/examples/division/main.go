package main

// gontract example using Require and Ensure

import (
	"fmt"

	"github.com/gontract/gontract"
	"github.com/obnoxxx/manc"
)

func Divide(dividend float64, divisor float64) (quotient float64) {

	// precondition:
	gontract.Require(divisor != 0, "divisor must be non-zero")
	// postcondition:
	defer func() {
		gontract.Ensure(manc.FloatsAreEqual(quotient*divisor, dividend), "quotient calculated correctly")
	}()

	quotient = dividend / divisor

	return
}

func main() {

	dividends := [...]float64{10.0, 4.0, 1.0, 0}
	var divisor = 2.0

	for _, dividend := range dividends {

		quotient := Divide(dividend, divisor)

		fmt.Printf(" %f divided by %f  is %f\n", dividend, divisor, quotient)
	}
}
