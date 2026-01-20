package main

import (
	"fmt"
	"math"

	"github.com/gontract/gontract"
)

func CalculateSqrt(n float64) float64 {

	gontract.PreCondition(n >= 0, "square root undefined for negative numbers")

	r := math.Sqrt(n)

	gontract.PostCondition(r*r == n, "input must equal square of result")

	return r
}

func main() {

	var num float64 = 1

	root := CalculateSqrt(num)

	fmt.Printf("square root of %f is %f\n", num, root)

	num = -1

	// bug: calling sqrt with negative number.
	// This violates the precondition, breaking the contract
	root = CalculateSqrt(num)

	fmt.Printf("square root of %f is %f\n", num, root)

}
