package main

import (
	// standard
	"encoding/json"
	"fmt"
	"net/http"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{}
	response["msg"] = "Feature not yet implemented"

	output, err := json.Marshal(response)
	if err != nil {
		elog.Println(err)
	}

	fmt.Fprintf(w, string(output))
}
