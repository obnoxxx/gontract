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

func CheckSquareRoot(num float64) (ret float64) {

	ret = 0

	defer catch_violation(&ret)
	ret = SquareRoot(num)
	return
}

func CheckSquare(num float64) (ret float64) {

	ret = 0

	defer catch_violation(&ret)
	ret = Square(num)

	return

}

func TestSquare(t *testing.T) {

	var num float64 = 1

	s := CheckSquare(num)

	//assert.True(s >= 0)
	assert.Equal(t, s, 1.0)

	num = 2
	s = Square(num)

	assert.Equal(t, s, 4.0)

	num = 3
	s = Square(num)

	assert.Equal(t, s, 9.0)

}

func TestSquareRoot(t *testing.T) {

	var num float64 = 1
	r := CheckSquareRoot(num)
	assert.Equal(t, r, 1.0)

	num = 4
	r = CheckSquareRoot(num)
	assert.Equal(t, r, 2.0)

	num = -1
	r = CheckSquareRoot(num)
	// on violation, -1 is returned:
	assert.Equal(t, r, -1.0)

}

func Test_main(t *testing.T) {
	main()
	// trivial assert as there is nothing to test
	assert.True(t, true)
}
