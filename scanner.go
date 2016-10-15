package main

import (
	// standard
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

// struct to hold compiler and channels
type Scanner struct {
	compiler  *yara.Compiler
	requests  chan string
	responses chan *Response
}

// load rule and process directory

func loadIndex(compiler *yara.Compiler, indexpath string) error {
	filehandle, err := os.Open(indexpath)
	if err != nil {
		return err
	}
	defer filehandle.Close()

	// generate namespace
	fields := strings.Split(indexpath, "/")
	filename := fields[len(fields)-1]
	namespace := strings.Split(filename, "_")[0]

	err = compiler.AddFile(filehandle, namespace)
	if err != nil {
		elog.Println(err)
		return err
	}

	return nil
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

func NewScanner(req chan string, resp chan *Response) (*Scanner, error) {
	scanner := new(Scanner)
	scanner.requests = req
	scanner.responses = resp

	compiler, err := yara.NewCompiler()
	if err != nil {
		return nil, err
	}
	scanner.compiler = compiler

	return scanner, nil
}

func (scanner *Scanner) LoadIndex(indexPath string) error {
	err := loadIndex(scanner.compiler, indexPath)
	if err != nil {
		elog.Println(err)
		return err
	}

	return nil
}

func (scanner *Scanner) Run() {
	info.Println("Waiting for scan requests")
	for request := range scanner.requests {
		response, err := scan(scanner.compiler, request)
		if err != nil {
			elog.Println(err)
		}
		scanner.responses <- response
	}
}
