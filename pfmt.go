package pfmt

import (
	"fmt"
	"io"
	"strings"
)

func Format(text string, rules Rule) string {
	return replacePlaceholders(text, rules, DEFAULT_START_MARKER, DEFAULT_END_MARKER)
}

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

func (f *Pfmt) WFormat(w io.Writer, text string) {
	fmt.Fprint(w, replacePlaceholders(text, f.rules, f.start_marker, f.end_marker))
}

func (f *Pfmt) PFormat(p *string, text string) {
	*p = f.Format(text)
}

func (f *Pfmt) SetRules(rules Rule) {
	f.rules = rules
}

func (f *Pfmt) AddRule(rules ...Rule) {
	f.rules = mergeRulesWithBase(f.rules, rules...)
}

func (f *Pfmt) Reset() {
	if len(f.rules) == 0 {
		return
	}
	f.rules = make(Rule)
}
