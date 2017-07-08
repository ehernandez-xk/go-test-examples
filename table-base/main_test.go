package main

import (
	"testing"
)

//TestReverse test reverse func
func TestReverse(t *testing.T) {
	table := []struct {
		input, output string
	}{
		{"hola", "aloh"},
		{"", ""},
	}

	for _, i := range table {
		got := reverse(i.input)
		if i.output != got {
			t.Errorf("Input: %s output: %s correct: %s", i.input, got, i.output)
		}
	}
}
