package main

import (
	// standard
	"fmt"
	"io/ioutil"
	"os"

	// external
	"github.com/hillu/go-yara"
	"github.com/jheise/yaramsg"
)

// struct to hold multiple matches
type Response struct {
	Filename string
	Matches  []*yaramsg.Match
}

// load yara rules
func loadRules(compilr *yara.Compiler) error {
	var retcode error
	filenames, err := ioutil.ReadDir(rules_dir)
	if err != nil {
		return err
	}

	for _, filename := range filenames {
		if !filename.IsDir() {
			fullpath := rules_dir + "/" + filename.Name()
			info.Println("Adding " + fullpath)

			filehandle, err := os.Open(fullpath)
			if err != nil {
				return err
			}
			defer filehandle.Close()

			err = compilr.AddFile(filehandle, "rules")
			if err != nil {
				fmt.Println(fullpath)
				fmt.Println(err)
			}
		}
	}

	return retcode
}

// load rules from file
func loadRulesFile(compilr *yara.Compiler) error {
	filehandle, err := os.Open(rules_dir)
	if err != nil {
		return err
	}
	defer filehandle.Close()

	err = compilr.AddFile(filehandle, "rules")
	return err
}

func scan(compiler *yara.Compiler, filename string) (*Response, error) {
	response := new(Response)
	response.Filename = filename
	var err error
	var matches []*yaramsg.Match

	rules, err := compiler.GetRules()
	if err != nil {
		return response, err
	}

	filepath := uploads_dir + "/" + filename
	output, err := rules.ScanFile(filepath, 0, 30)
	if err != nil {
		return response, err
	}

	for _, resp := range output {
		match := new(yaramsg.Match)
		match.Rule = resp.Rule
		match.Namespace = resp.Namespace
		match.Tags = resp.Tags

		matches = append(matches, match)

	}

	response.Matches = matches

	return response, err

}

// create a compiler, read in rules, wait for requests to scan

func scanner() {
	info.Println("Creating compiler")
	compiler, err := yara.NewCompiler()
	if err != nil {
		elog.Println(err)
	}

	info.Println("Loading rules from " + rules_dir)
	err = loadRules(compiler)
	//err = loadRulesFile(compiler)
	if err != nil {
		elog.Println(err)
	}

	info.Println("Waiting for scan requests")
	for request := range requests {
		response, err := scan(compiler, request)
		if err != nil {
			elog.Println(err)
		}
		responses <- response
	}
}
