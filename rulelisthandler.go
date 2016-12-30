package main

import (
	// standard
	"encoding/json"
	"fmt"
	"net/http"
	// external
	"github.com/gorilla/mux"
	//"github.com/jheise/yaramsg"
)

func RuleListHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ruleset := vars["ruleset"]

	req := NewRuleListRequest(ruleset)
	rulerequests<- req

	response := <-req.ResponseChan

	output, err := json.Marshal(response)
	if err != nil {
		elog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, string(output))
}
