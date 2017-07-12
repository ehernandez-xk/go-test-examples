package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacto(t *testing.T) {
	t1 := 5 + 4 + 3 + 2 + 1
	t2 := facto(5)
	assert.Equal(t, t1, t2, "they should be equal")

	f := facto(-1)
	assert.Equal(t, 0, f, "factorial of -1 should be 0")
}

//I'm testing the Atoi(string) function from the package strconv
func TestAtoi(t *testing.T) {
	tests := map[string]struct {
		input  string
		output int
		err    error
	}{
		"successful conversion": {input: "1", output: 1, err: nil},
		"invalid integer":       {input: "a text", output: 0, err: &strconv.NumError{}},
	}

	for testName, test := range tests {
		t.Logf("Running test case %s", testName)

		output, err := strconv.Atoi(test.input)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.output, output)
	}
}
