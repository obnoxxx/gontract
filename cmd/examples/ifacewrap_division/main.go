package main

import (
	"fmt"

	"github.com/gontract/gontract/ifacewrap"
	"github.com/obnoxxx/manc"
)

type Divider interface {
	Divide(dividend, divisor float64) (quotient float64)
}

type divider struct{}

func (divider) Divide(dividend, divisor float64) float64 {
	return dividend / divisor
}

type divideInput struct {
	dividend float64
	divisor  float64
}

type divideOutput struct {
	quotient float64
}

type contractDivider struct {
	inner    Divider
	contract ifacewrap.Contract[divideInput, divideOutput]
}

func newContractDivider(inner Divider) Divider {
	return contractDivider{
		inner: inner,
		contract: ifacewrap.Contract[divideInput, divideOutput]{
			Requirements: []ifacewrap.Requirement[divideInput]{
				{
					Predicate: func(in divideInput) bool {
						return !manc.FloatIsZero(in.divisor)
					},
					Message: "divisor must be non-zero",
				},
			},
			Assurances: []ifacewrap.Assurance[divideInput, divideOutput]{
				{
					Predicate: func(in divideInput, out divideOutput) bool {
						return manc.FloatsAreEqual(out.quotient*in.divisor, in.dividend)
					},
					Message: "quotient calculated correctly",
				},
			},
		},
	}
}

func (d contractDivider) Divide(dividend, divisor float64) float64 {
	out := ifacewrap.Call(
		divideInput{dividend: dividend, divisor: divisor},
		func() divideOutput {
			return divideOutput{
				quotient: d.inner.Divide(dividend, divisor),
			}
		},
		d.contract,
	)

	return out.quotient
}

func main() {
	dividends := [...]float64{10.0, 4.0, 1.0, 0}
	d := newContractDivider(divider{})
	divisor := 2.0

	for _, dividend := range dividends {
		quotient := d.Divide(dividend, divisor)
		fmt.Printf(" %f divided by %f  is %f\n", dividend, divisor, quotient)
	}
}
