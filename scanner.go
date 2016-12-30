package main

import (
	// standard
	"fmt"

	// external
	"github.com/jheise/yaramsg"
)

// struct to hold compiler and channels
type Scanner struct {
	rulesets     []*RuleSet
	scanrequests chan *ScanRequest
	namerequests chan *RuleSetRequest
	rulerequests chan *RuleListRequest
}

func (self *Scanner) scan(filename string) (*ScanResponse, error) {
	response := new(ScanResponse)
	response.Filename = filename
	var err error
	var matches []*yaramsg.Match

	filepath := uploads_dir + "/" + filename

	for _, ruleset := range self.rulesets {
		output, err := ruleset.Rules.ScanFile(filepath, 0, 300)
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

	}
	response.Matches = matches

	return response, err
}

func (self *Scanner) LoadIndex(indexPath string) error {
	ruleset, err := NewRuleSet(indexPath)
	if err != nil {
		return err
	}
	self.rulesets = append(self.rulesets, ruleset)
	return nil
}

func (self *Scanner) listRuleSets() *RuleSetResponse {
	response := new(RuleSetResponse)
	for _, ruleset := range self.rulesets {
		response.Names = append(response.Names, ruleset.Name)
	}

	return response

}

func (self *Scanner) listRules(rulesetname string) (*RuleListResponse, error) {
	response := new(RuleListResponse)
	fmt.Printf("listRules called, %s\n", rulesetname)
	for _, ruleset := range self.rulesets {
		fmt.Printf("Looking for %s, looking at %s\n", rulesetname, ruleset.Name)
		if ruleset.Name == rulesetname {
			rules, err := ruleset.ListRules()
			if err != nil {
				return nil, err
			}
			for _, rule := range rules {
				response.Rules = append(response.Rules, rule)
			}
		}
	}

	return response, nil
}

func (self *Scanner) Run() {
	info.Println("Waiting for scan requests")
	for {
		select {
		case scanmsg := <-scanrequests:
			response, err := self.scan(scanmsg.Filename)
			if err != nil {
				elog.Println(err)
				return
			}
			scanmsg.ResponseChan <- response
		case setmsg := <-namerequests:
			response := self.listRuleSets()
			setmsg.ResponseChan <- response
		case rulemsg := <-rulerequests:
			response, err := self.listRules(rulemsg.RuleSet)
			if err != nil {
				elog.Println(err)
				return
			}
			rulemsg.ResponseChan <- response
		}
	}
}

func NewScanner(scan chan *ScanRequest, name chan *RuleSetRequest, list chan *RuleListRequest) (*Scanner, error) {
	scanner := new(Scanner)
	scanner.scanrequests = scan
	scanner.namerequests = name
	scanner.rulerequests = list

	return scanner, nil
}
