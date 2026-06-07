package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func catch_violation(num *float64) {

	if r := recover(); r != nil {
		*num = -1.0
	}
}

func checkDivide(dividend float64, divisor float64) (ret float64) {

	ret = 0
	defer catch_violation(&ret)

	ret = Divide(dividend, divisor)
	return

}

func TestDivide(t *testing.T) {

	cases := [...]struct {
		dividend float64
		divisor  float64
		quotient float64
	}{
		{10.0, 2.0, 5.0},
		{4.0, 2.0, 2.0},
		{0.0, 2.0, 0.0},
		{1.0, 0.0, -1.0},
	}
	for _, c := range cases {
		quotient := checkDivide(c.dividend, c.divisor)
		assert.Equal(t, quotient, c.quotient)
	}
}

func Test_main(t *testing.T) {
	var err error = nil
	main()

	// trivial condition as main has no result
	assert.Equal(t, err, nil)

}
