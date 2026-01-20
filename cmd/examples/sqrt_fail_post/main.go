package main

import (
	"fmt"

	"github.com/gontract/gontract"
)

func CalculateSqrt(n float64) float64 {

	gontract.PreCondition(n >= 0, "square root undefined for negative numbers")

	// Bug: This calculation of the result is wrong and breaks
	// contract (see postcondition below).
	r := n * n

	gontract.PostCondition(r*r == n, "input must equal square of result")

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
