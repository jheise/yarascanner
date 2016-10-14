package main

import (
	// standard
	"fmt"
	"io/ioutil"
	"net/http"

	// external
	"github.com/gorilla/mux"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	info.Printf("Serving %s\n", filename)

	filepath := fullpath(filename)

	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		elog.Println(err)
	}

	fmt.Fprintf(w, string(data))
}
