package main

import (
	"fmt"
	"math"

	"github.com/gontract/gontract"
	"github.com/obnoxxx/manc"
)

func CalculateSqrt(n float64) (r float64) {

	// PRECONDITIONS:
	gontract.Require(!math.IsNaN(n), "Input must be a number.")
	gontract.Require(!math.IsInf(n, 0), "Input must be finite.")
	gontract.Require(n >= 0, "Input must be non-negative.")
	// POSTCONDITIONS:
	defer func() {
		gontract.Ensure(!math.IsNaN(r), "result must be a number.")
		gontract.Ensure(manc.FloatsAreEqual(r*r, n), "Square of result must equal input.")
	}()

	r = math.Sqrt(n)

	return
}

func main() {

	var num float64 = 1

	root := CalculateSqrt(num)

	fmt.Printf("The square root of %f is %f\n", num, root)

	num = 4

	root = CalculateSqrt(num)

	fmt.Printf("The square root of %f is %f\n", num, root)

}
