package gontract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkCondition(predicate bool, kind Kind, msg string) (ret string) {
	ret = "foo"
	defer func() {
		if r := recover(); r != nil {
			ret = r.(string)
		}
	}()

	Condition(predicate, kind, msg)
	return

}

func TestCondition(t *testing.T) {
	cases := []struct {
		predicate bool
		kind      Kind
		msg       string
		expected  string
	}{
		{true, KindPre, "trivially true", "foo"},
		{false, KindPost, "trivially false", "postcondition not satisfied (trivially false) - software bug!?"},
	}

	for _, c := range cases {

		ret := checkCondition(c.predicate, c.kind, c.msg)
		assert.Equal(t, c.expected, ret)

	}

}

func CheckPreCondition(predicate bool, msg string) (ret string) {
	ret = "foo"
	defer func() {
		if r := recover(); r != nil {
			ret = r.(string)
		}
	}()

	PreCondition(predicate, msg)
	return
}
func TestPreCondition(t *testing.T) {

	ret := CheckPreCondition(true, "trivially true")

	assert.Equal(t, ret, "foo")

	ret = CheckPreCondition(false, "trivially false")

	assert.Equal(t, ret, "precondition not satisfied (trivially false) - software bug!?")
}
func CheckPostCondition(predicate bool, msg string) (ret string) {
	ret = "foo"
	defer func() {
		if r := recover(); r != nil {
			ret = r.(string)
		}
	}()

	PostCondition(predicate, msg)
	return
}
func TestPostCondition(t *testing.T) {

	ret := CheckPostCondition(true, "trivially true")
	assert.Equal(t, ret, "foo")

	ret = CheckPostCondition(false, "trivially false")
	assert.Equal(t, ret, "postcondition not satisfied (trivially false) - software bug!?")
}
