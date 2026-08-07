package main

import (
	"fmt"

	"github.com/gontract/gontract"
	"github.com/obnoxxx/manc"
)

func CalculateSqrt(n float64) (r float64) {

	// precondition:
	gontract.Require(n >= 0, "square root undefined for negative numbers")
	// postcondition:
	defer gontract.Ensure(manc.FloatsAreEqual(r*r, n), "square root calculated correctly")

	// Bug: This calculation of the result is wrong and breaks
	// contract (see postcondition below).
	r = n * n

	return r
}

func main() {

	var num float64 = 1

	root := CalculateSqrt(num)

	fmt.Printf("The square root of %f is %f\n", num, root)

	num = 4

	root = CalculateSqrt(num)

	fmt.Printf("The square root of %f is %f\n", num, root)

}
