package main

import (
	// standard
	"bufio"
	"fmt"
	"os"
	"strings"

	// external
	"github.com/hillu/go-yara"
)

type RuleSet struct {
	Name     string
	FilePath string
	Compiler *yara.Compiler
	Rules    *yara.Rules
}

func (self *RuleSet) ListRules() ([]string, error) {
	rules := []string{}
	fmt.Printf("ListRules called")

	file, err := os.Open(self.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewScanner(file)
	for reader.Scan() {
		rules = append(rules, reader.Text())
	}

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return rules, nil
}

func NewRuleSet(indexpath string) (*RuleSet, error) {
	filehandle, err := os.Open(indexpath)
	if err != nil {
		return nil, err
	}
	defer filehandle.Close()

	fields := strings.Split(indexpath, "/")
	filename := fields[len(fields)-1]
	namespacestr := strings.Split(filename, "_")[0]

	compiler, err := yara.NewCompiler()
	if err != nil {
		return nil, err
	}

	err = compiler.AddFile(filehandle, namespacestr)
	if err != nil {
		elog.Println(err)
		return nil, err
	}

	rules, err := compiler.GetRules()
	if err != nil {
		return nil, err
	}

	namespace := new(RuleSet)
	namespace.FilePath = indexpath
	namespace.Name = namespacestr
	namespace.Compiler = compiler
	namespace.Rules = rules

	return namespace, nil
}
