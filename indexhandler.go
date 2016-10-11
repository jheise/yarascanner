package main

import (
	// standard
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>HELLO</title></head><body>/list<br/></body></html>")
}
