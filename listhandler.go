package main

import (
	// standard
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	// external
	"github.com/jheise/yaramsg"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	files := new(yaramsg.ListResponse)
	filenames, err := ioutil.ReadDir(uploads_dir)
	if err != nil {
		elog.Println(err)
	}

	for _, filename := range filenames {
		files.Files = append(files.Files, filename.Name())
	}

	output, err := json.Marshal(files)
	if err != nil {
		elog.Println(err)
	}

	fmt.Fprintf(w, string(output))
}
