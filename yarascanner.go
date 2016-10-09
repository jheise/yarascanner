package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	//external
	"github.com/hillu/go-yara"
)

var (
	source_dir  string
	target_file string
)

func init() {
	flag.StringVar(&source_dir, "rules", "rules", "path to yara rules")
	flag.StringVar(&target_file, "target", "target", "path to target file")
	flag.Parse()
}

func loadRules(compilr *yara.Compiler) error {
	var retcode error
	filenames, err := ioutil.ReadDir(source_dir)
	if err != nil {
		return err
	}

	for _, filename := range filenames {
		if !filename.IsDir() {
			fullpath := source_dir + "/" + filename.Name()
			fmt.Println("Adding " + fullpath)

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

func main() {
	fmt.Printf("Loading rules from %s\n", source_dir)
	fmt.Printf("generating yara compiler")
	complr, err := yara.NewCompiler()

	if err != nil {
		panic(err)
	}
	fmt.Printf("complr %s\n", complr)

	err = loadRules(complr)
	if err != nil {
		panic(err)
	}

	rules, err := complr.GetRules()
	if err != nil {
		panic(err)
	}

	fmt.Println("Scanning file: " + target_file)
	output, err := rules.ScanFile(target_file, 0, 30)
	if err != nil {
		panic(err)
	}
	for _, resp := range output {
		fmt.Println("Matched rule " + resp.Rule)
	}

}
