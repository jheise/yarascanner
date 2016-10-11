package main

import (
	// standard
	"encoding/json"
	"fmt"
	"net/http"

	// external
	"github.com/jheise/yaramsg"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	response := new(yaramsg.UploadResponse)
	response.Message = "Feature not yet implemented"
	response.Result = false

	output, err := json.Marshal(response)
	if err != nil {
		elog.Println(err)
	}

	fmt.Fprintf(w, string(output))
}
