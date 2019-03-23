package greek

import (
	"regexp"
)

type rule struct {
	expr        *regexp.Regexp
	replacement string
}

func applyRules(s string, rules []rule) string {
	res := []byte(s)

	for _, rule := range rules {
		res = rule.expr.ReplaceAll(res, []byte(rule.replacement))
	}

	return string(res)
}
