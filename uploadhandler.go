package main

import (
	// standard
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// external
	"github.com/gorilla/mux"
	"github.com/jheise/yaramsg"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	info.Printf("Uploading %s\n", filename)

	response := new(yaramsg.UploadResponse)
	response.Message = "Successfully uploaded " + filename
	response.Result = true

	buf, _ := ioutil.ReadAll(r.Body)
	filepath := fullpath(filename)
	err := ioutil.WriteFile(filepath, buf, 0644)
	if err != nil {
		elog.Println(err)
		response.Message = "Failed to upload " + filename
		response.Result = false
	}

	output, err := json.Marshal(response)
	if err != nil {
		elog.Println(err)
	}

	fmt.Fprintf(w, string(output))
}
