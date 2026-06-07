package main

import (
	"fmt"
	"math"

	"github.com/gontract/gontract/funcwrap"
)

type divideInput struct {
	dividend float64
	divisor  float64
}
type divideOutput struct {
	quotient float64
}

const epsilon = 1e-8

func floatEquals(a, b float64) bool {
	diff := math.Abs(a - b)
	if diff < epsilon {
		return true
	}

	largest := math.Max(math.Abs(a), math.Abs(b))
	return diff < largest*epsilon
}

func Divide(dividend float64, divisor float64) (quotient float64) {
	quotient = funcwrap.Call(
		divideInput{dividend: dividend, divisor: divisor},
		func() divideOutput {
			return divideOutput{quotient: dividend / divisor}
		},
		funcwrap.Contract[divideInput, divideOutput]{
			Requirements: []funcwrap.Requirement[divideInput]{
				{
					Predicate: func(in divideInput) bool { return in.divisor != 0 },
					Message:   "divisor must be non-zero",
				},
			},
			Assurances: []funcwrap.Assurance[divideInput, divideOutput]{
				{
					Predicate: func(in divideInput, out divideOutput) bool {
						return floatEquals(out.quotient*in.divisor, in.dividend)
					},
					Message: "quotient calculated correctly",
				},
			},
		},
	).quotient

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
