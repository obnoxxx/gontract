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

	fmt.Printf("The square root of %f is %f\n", num, root)

	num = 4

	root = CalculateSqrt(num)

	fmt.Printf("The square root of %f is %f\n", num, root)

}
