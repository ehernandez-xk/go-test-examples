package main

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace.")
		if buf.String() != "Hello trace.\n" {
			t.Errorf("Trace should not write '%s'.", buf.String())
		}
	}
}
