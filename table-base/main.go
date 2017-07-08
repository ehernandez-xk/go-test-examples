package main

import (
	"fmt"
	"strings"
)

func reverse(tx string) (out string) {
	s := strings.Split(tx, "")
	for i := len(s) - 1; i != -1; i-- {
		out = out + s[i]
	}
	return
}

func main() {
	fmt.Println(reverse("hola"))
}
