package pfmt

import "strings"

func (f *Pfmt) SetMarkers(start string, end string) {
	s := strings.TrimSpace(start)
	e := strings.TrimSpace(end)

	if len(s) == 0 || len(e) == 0 {
		return
	}

	f.start_marker = s
	f.end_marker = e
}

func (f *Pfmt) Format(text string) string {
	return replacePlaceholders(text, f.rules, f.start_marker, f.end_marker)
}

func (f *Pfmt) SetRules(rules map[string]interface{}) {
	f.rules = rules
}

func (f *Pfmt) AddRule(rules ...map[string]interface{}) {
	f.rules = mergeMapsWithBase(f.rules, rules...)
}

func (f *Pfmt) Reset() {
	f.rules = make(map[string]interface{})
}