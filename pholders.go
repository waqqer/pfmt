package pfmt

import (
	"strings"
	"fmt"
)

func replacePlaceholders(text string, rules map[string]interface{}, start_marker string, end_marker string) string {
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