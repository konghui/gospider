package filter

import "regexp"

type FilterRule struct {
	Rule     string
	Parttern *regexp.Regexp
	Data     string
}

func NewFilterRule(rule string) (filterRule *FilterRule) {
	filterRule = &FilterRule{Rule: rule, Parttern: regexp.MustCompile(rule)}
	return
}
