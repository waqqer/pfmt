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
	return f.replaceAll(text)
}

func (f *Formatter) SetRules(rules map[string]interface{}) {
	f.rules = rules
}

func (f *Formatter) replaceAll(text string) string {
	if len(f.rules) == 0 {
		return text
	}

	var result strings.Builder
	for len(text) > 0 {
		start := strings.Index(text, f.start_marker)
		if start == -1 {
			result.WriteString(text)
			break
		}

		result.WriteString(text[:start])
		text = text[start+len(f.start_marker):]

		end := strings.Index(text, f.end_marker)
		if end == -1 {
			result.WriteString(f.start_marker)
			result.WriteString(text)
			break
		}

		key := text[:end]
		value, exists := f.rules[key]
		if !exists {
			result.WriteString(f.start_marker)
			result.WriteString(key)
			result.WriteString(f.end_marker)
		} else if fn, ok := value.(func() string); ok {
			result.WriteString(fn())
		} else if fn, ok := value.(func() interface{}); ok {
			result.WriteString(fmt.Sprintf("%v", fn()))
		} else {
			result.WriteString(fmt.Sprint(value))
		}

		text = text[end+len(f.end_marker):]
	}

	return result.String()
}