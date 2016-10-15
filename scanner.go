package main

import (
	// standard
	"io/ioutil"
	"os"
	"strings"

	// external
	"github.com/hillu/go-yara"
	"github.com/jheise/yaramsg"
)

// struct to hold multiple matches
type Response struct {
	Filename string
	Matches  []*yaramsg.Match
}

// load rule and process directory

func loadRule(compiler *yara.Compiler, namespace string, rulepath string) error {
	var retval error

	filehandle, err := os.Open(rulepath)
	if err != nil {
		return err
	}
	defer filehandle.Close()

	err = compiler.AddFile(filehandle, namespace)
	if err != nil {
		elog.Println(err)
	}
	return retval
}

// load all indexes in a directory

func processDirectory(compiler *yara.Compiler, rules_src string) {
	filenames, err := ioutil.ReadDir(rules_src)
	if err != nil {
		panic(err)
	}

	for _, filename := range filenames {
		fullpath := rules_src + "/" + filename.Name()
		if strings.HasSuffix(filename.Name(), ".yar") && filename.Name() != "index.yar" {
			info.Println("Loading Rule " + filename.Name())
			namespace := strings.Split(filename.Name(), "_")[0]
			err = loadRule(compiler, namespace, fullpath)
			if err != nil {
				panic(err)
			}
		}
	}
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
	output, err := rules.ScanFile(filepath, 0, 300)
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
	//err = loadRules(compiler)
	//err = loadRulesFile(compiler)
	//if err != nil {
	//	elog.Println(err)
	//}
	processDirectory(compiler, rules_dir)

	info.Println("Waiting for scan requests")
	for request := range requests {
		response, err := scan(compiler, request)
		if err != nil {
			elog.Println(err)
		}
		responses <- response
	}
}
