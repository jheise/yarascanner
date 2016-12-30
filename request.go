package main

// struct to make requests to the scanner
type ScanRequest struct {
	Filename     string
	ResponseChan chan *ScanResponse
}

func NewScanRequest(filename string) *ScanRequest {
	scan := new(ScanRequest)
	scan.Filename = filename
	scan.ResponseChan = make(chan *ScanResponse)

	return scan
}

type RuleSetRequest struct {
	ResponseChan chan *RuleSetResponse
}

func NewRuleSetRequest() *RuleSetRequest {
	rule := new(RuleSetRequest)
	rule.ResponseChan = make(chan *RuleSetResponse)

	return rule
}

type RuleListRequest struct {
	RuleSet      string
	ResponseChan chan *RuleListResponse
}

func NewRuleListRequest(ruleset string) *RuleListRequest {
	rule := new(RuleListRequest)
	rule.RuleSet = ruleset
	rule.ResponseChan = make(chan *RuleListResponse)

	return rule
}
