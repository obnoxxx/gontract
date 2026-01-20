package main

import (
	"testing"

	"github.com/gontract/gontract"
	"github.com/stretchr/testify/assert"
)

func catch_violation(num *float64) {

	var str = "foo"

	gontract.CatchViolation(&str)
	if str != "foo" {
		*num = -1.0
	}

}

func checkCalculateSqrt(num float64) (ret float64) {

	ret = 0
	defer catch_violation(&ret)

	ret = CalculateSqrt(num)
	return

}

func TestCalculateSqrt(t *testing.T) {

	cases := [...]struct {
		num  float64
		root float64
	}{
		{4.0, -1.0},
		{9.0, -1.0},
		{-2.0, -1.0},
	}
	for _, c := range cases {
		ret := checkCalculateSqrt(c.num)
		assert.Equal(t, ret, c.root)

	}
}

func check_main() (ret float64) {
	ret = 0
	defer catch_violation(&ret)

	main()
	return
}

func Test_main(t *testing.T) {
	//var ret float64
	ret := 0.0

	ret = check_main()
	// trivial condition as main has no result
	assert.Equal(t, ret, -1.0)

}
