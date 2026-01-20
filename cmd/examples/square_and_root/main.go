package main

import (
	"fmt"
	"math"

	"github.com/gontract/gontract"
)

func Square(n float64) (s float64) {

	gontract.Require(true, "no special precondition")

	s = n * n

	gontract.Ensure(s >= 0, "square of any number is non-negative")

	return

}

func SquareRoot(n float64) (r float64) {

	gontract.Require(n >= 0, "square root undefined for negative numbers")

	r = math.Sqrt(n)

	gontract.Ensure(Square(r) == n, "input must equal square of result")

	return

}

func main() {

	var num float64 = 1

	root := SquareRoot(num)

	fmt.Printf("The square root of %f is %f\n", num, root)

	s := Square(num)

	fmt.Printf("The square of %f is %f\n", num, s)

	num = 4

	root = SquareRoot(num)

	fmt.Printf("The square root of %f is %f\n", num, root)
	s = Square(num)

	fmt.Printf("The square of %f is %f\n", num, s)

	num = 9

	root = SquareRoot(num)

	fmt.Printf("The square root of %f is %f\n", num, root)

	s = Square(num)

	fmt.Printf("The square of %f is %f\n", num, s)
}
