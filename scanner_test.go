package main

import (
	"testing"
)

func buildChannels() (chan string, chan *Response) {
	requests := make(chan string)
	responses := make(chan *Response)

	return requests, responses
}

func TestNewScanner(t *testing.T) {
	requests, responses := buildChannels()

	scanner, err := NewScanner(requests, responses)
	if err != nil {
		t.Error(err)
	}
	if scanner.compiler == nil {
		t.Error("Compiler should not be nil")
	}
}

func TestLoadIndex(t *testing.T) {
	requests, responses := buildChannels()

	scanner, err := NewScanner(requests, responses)
	if err != nil {
		t.Error(err)
	}

	err = scanner.LoadIndex("testdata/rules/malware_index.yar")
	if err != nil {
		t.Error(err)
	}
}
