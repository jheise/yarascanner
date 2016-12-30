package main

import (
	// standard
	"encoding/json"
	"fmt"
	"net/http"
	// external
	//"github.com/jheise/yaramsg"
)

func RuleSetListHandler(w http.ResponseWriter, r *http.Request) {
	req := NewRuleSetRequest()
	namerequests <- req

	response := <-req.ResponseChan

	output, err := json.Marshal(response)
	if err != nil {
		elog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(output))
}
