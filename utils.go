package main

import (
	// standard
	"fmt"
	"io/ioutil"
)

type StringArgs []string

func (s *StringArgs) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *StringArgs) Set(value string) error {
	*s = append(*s, value)

	return nil
}

func fileExists(target string) (bool, error) {
	retval := false
	var err error

	filenames, err := ioutil.ReadDir(uploads_dir)
	if err != nil {
		return retval, err
	}

	for _, filename := range filenames {
		if filename.Name() == target {
			retval = true
			return retval, err
		}
	}
	return retval, err
}

func fullpath(filepath string) string {
	return uploads_dir + "/" + filepath
}
