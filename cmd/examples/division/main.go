package main

// gontract example using Require and Ensure

import (
	"fmt"
	"math"

	"github.com/gontract/gontract"
)

const epsilon = 1e-8

// floatEquals():
// function to use instead of == for comparing floats.
// This takes rounding errors into account.
func floatEquals(a, b float64) bool {
	// absolute comparison:
	diff := math.Abs(a - b)
	if diff < epsilon {
		return true
	}
	// additional relative comparison to avoid false negatives
	largest := math.Max(math.Abs(a), math.Abs(b))
	return diff < largest*epsilon
}

func Divide(dividend float64, divisor float64) (quotient float64) {

	// precondition:
	gontract.Require(divisor != 0, "divisor must be non-zero")
	// postcondition:
	defer func() {
		gontract.Ensure(floatEquals(quotient*divisor, dividend), "quotient calculated correctly")
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
