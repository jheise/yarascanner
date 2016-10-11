package main

import (
	// standard
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	// external
	"github.com/gorilla/mux"
	"github.com/jheise/yaramsg"
)

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	info.Printf("Removing %s\n", filename)

	// check that file exists with traversal safe function
	fileexists, err := fileExists(filename)
	if err != nil {
		return
	}

	if fileexists != true {
		return
	}

	response := new(yaramsg.RemoveResponse)
	response.Message = "Removing " + filename
	err = os.Remove(fullpath(filename))
	if err != nil {
		response.Result = false
	} else {
		response.Result = true
	}

	output, err := json.Marshal(response)
	if err != nil {
		elog.Println(err)
	}

	fmt.Fprintf(w, string(output))
}
