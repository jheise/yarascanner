package main

import (
	// standard
	"encoding/json"
	"fmt"
	"net/http"

	// external
	"github.com/gorilla/mux"
)

func ScanHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	info.Printf("Scanning %s\n", filename)

	// check that file exists with traversal safe function
	fileexists, err := fileExists(filename)
	if err != nil {
		return
	}

	if fileexists != true {
		return
	}

	// send request for scanning
	requests <- filename

	// wait for response, need to be fixed
	var response *Response
	for current := range responses {
		if current.Filename == filename {
			response = current
			break
		}
	}

	output, err := json.Marshal(response)
	if err != nil {
		elog.Println(err)
	}

	fmt.Fprintf(w, string(output))
}
