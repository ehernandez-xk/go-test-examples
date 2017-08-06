package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/name/", nameHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// nameHandler receives a name in the path and says Hi to that name
func nameHandler(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	name := args[2]
	fmt.Fprintf(w, "Hi %s", name)
}
