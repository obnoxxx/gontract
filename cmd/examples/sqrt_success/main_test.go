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

func CheckCalculateSqrt(num float64) (ret float64) {

	ret = 0

	defer catch_violation(&ret)
	/*
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recovered from panic(CalculateSqrt):  '%s'", r.(string))
				ret = -1
			}
		}()
	*/
	ret = CalculateSqrt(num)
	return
}

func TestCalculateSqrt(t *testing.T) {

	var num float64 = 1
	root := CheckCalculateSqrt(num)
	assert.Equal(t, num, root*root)

	num = 4
	root = CheckCalculateSqrt(num)
	assert.Equal(t, num, root*root)

	num = -1
	root = CheckCalculateSqrt(num)
	// on violation, -1 is returned:
	assert.Equal(t, root, -1.0)
}

func Test_main(t *testing.T) {
	main()
	// trivial assert as there is nothing to test
	assert.Equal(t, true, true)
}
