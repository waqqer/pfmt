package pfmt

import (
	"fmt"
	"strings"
)

const (
	start_marker string = "%"
	end_marker string = "%"
)

type Formatter struct {
	rules  map[string]interface{}
}

func NewFormatter(rules map[string]interface{}) *Formatter {
	return &Formatter{
		rules:  rules,
	}
}

func (f *Formatter) Format(text string) string {
	return f.replaceAll(text)
}

func (f *Formatter) replaceAll(text string) string {
	if len(f.rules) == 0 {
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
		value, exists := f.rules[key]
		if !exists {
			result.WriteString(start_marker)
			result.WriteString(key)
			result.WriteString(end_marker)
		} else if fn, ok := value.(func() string); ok {
			result.WriteString(fn())
		} else {
			result.WriteString(fmt.Sprint(value))
		}

		text = text[end+len(end_marker):]
	}

	return result.String()
}
