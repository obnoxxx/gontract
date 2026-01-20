package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	main()
	// trivial assert as there is nothing to test
	assert.Equal(t, true, true)
}
