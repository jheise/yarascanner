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
		elog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if fileexists != true {
		http.Error(w, fmt.Sprintf("%s does not exist", filename), http.StatusInternalServerError)
		return
	}

	// send request for scanning
	newRequest := NewScanRequest(filename)
	scanrequests <- newRequest

	response := <-newRequest.ResponseChan

	output, err := json.Marshal(response)
	if err != nil {
		elog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(output))
}
