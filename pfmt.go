package pfmt

import (
	"fmt"
	"strings"
)

type Formatter struct {
	rules  map[string]interface{}
	start_marker string
	end_marker string
}

func NewEmptyFormatter() *Formatter {
	return &Formatter{
		rules: map[string]interface{}{},
		start_marker: "%",
		end_marker: "%",
	}
}

func NewFormatter(rules map[string]interface{}) *Formatter {
	return &Formatter{
		rules: rules,
		start_marker: "%",
		end_marker: "%",
	}
}

func (f *Formatter) SetMarkers(start string, end string) {
	f.start_marker = strings.TrimSpace(start)
	f.end_marker = strings.TrimSpace(end)
}

func (f *Formatter) Format(text string) string {
	return replacePlaceholders(text, f.rules, f.start_marker, f.end_marker)
}

func (f *Formatter) SetRules(rules map[string]interface{}) {
	f.rules = rules
}

func replacePlaceholders(text string, rules map[string]interface{}, start_marker string, end_marker string) string {
	start_marker = strings.TrimSpace(start_marker)
	end_marker = strings.TrimSpace(end_marker)

	if len(rules) == 0 {
		return text
	}

	var result strings.Builder
	for len(text) > 0 {
		start := strings.Index(text, start_marker)
		if start == -1 {
			result.WriteString(text)
			break
		}

		result.WriteString(text[:start])
		text = text[start+len(start_marker):]

		end := strings.Index(text, end_marker)
		if end == -1 {
			result.WriteString(start_marker)
			result.WriteString(text)
			break
		}

		key := text[:end]
		value, exists := rules[key]
		if !exists {
			result.WriteString(start_marker)
			result.WriteString(key)
			result.WriteString(end_marker)
		} else if fn, ok := value.(func() string); ok {
			result.WriteString(fn())
		} else if fn, ok := value.(func() interface{}); ok {
			fmt.Fprintf(&result, "%v", fn())
		} else {
			fmt.Fprint(&result, value)
		}

		text = text[end+len(end_marker):]
	}

	return result.String()
}