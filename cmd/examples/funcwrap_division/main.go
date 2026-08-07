package main

import (
	"fmt"

	"github.com/gontract/gontract/funcwrap"
	"github.com/obnoxxx/manc"
)

type divideInput struct {
	dividend float64
	divisor  float64
}
type divideOutput struct {
	quotient float64
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
					Predicate: func(in divideInput) bool { return !manc.FloatIsZero(in.divisor) },
					Message:   "divisor must be non-zero",
				},
			},
			Assurances: []funcwrap.Assurance[divideInput, divideOutput]{
				{
					Predicate: func(in divideInput, out divideOutput) bool {
						return manc.FloatsAreEqual(out.quotient*in.divisor, in.dividend)
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
