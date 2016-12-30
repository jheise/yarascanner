package main

import (
	// external
	"github.com/jheise/yaramsg"
)

// struct to hold multiple matches
type ScanResponse struct {
	Filename string
	Matches  []*yaramsg.Match
}

// struct to handle namespace requests
type RuleSetResponse struct {
	Names []string
}

// sturc to handle
type RuleListResponse struct {
	Rules []string
}
