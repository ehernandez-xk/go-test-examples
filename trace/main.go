package main

import (
	"fmt"
	"io"
	"os"
)

// Tracer is the interface that describes an object capable of
// tracing events throughout code
type Tracer interface {
	Trace(...interface{})
}

//New returns a new tracer
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

// prints in the stdout a text
func main() {
	t := New(os.Stdout)
	t.Trace("hola mundo")
}
