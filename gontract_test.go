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
