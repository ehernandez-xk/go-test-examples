package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	t1 := 5 + 4 + 3 + 2 + 1
	t2 := facto(5)
	assert.Equal(t, t1, t2, "they should be equal")

	f := facto(-1)
	assert.Equal(t, 0, f, "factorial of -1 should be 0")
}
