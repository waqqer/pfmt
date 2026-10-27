package pfmt

import "testing"

func TestFormatterFormat(t *testing.T) {
	tests := []struct {
		name  string
		rules map[string]interface{}
		input string
		want  string
	}{
		{
			name:  "replaces string value",
			rules: map[string]interface{}{"name": "Alice"},
			input: "Hello, %name%!",
			want:  "Hello, Alice!",
		},
		{
			name:  "replaces numeric value",
			rules: map[string]interface{}{"number": 123},
			input: "Number: %number%",
			want:  "Number: 123",
		},
		{
			name: "calls function value",
			rules: map[string]interface{}{
				"func": func() string { return "hello world" },
			},
			input: "Result: %func%",
			want:  "Result: hello world",
		},
		{
			name:  "keeps unknown placeholder",
			rules: map[string]interface{}{"text": "12:00:00"},
			input: "Date: %date%",
			want:  "Date: %date%",
		},
		{
			name: "replaces known placeholders and keeps unknown ones",
			rules: map[string]interface{}{
				"text":   "12:00:00",
				"number": 123,
				"func":   func() string { return "hello world" },
			},
			input: "Text: %text% | Num: %number% | Func: %func% | Date: %date%",
			want:  "Text: 12:00:00 | Num: 123 | Func: hello world | Date: %date%",
		},
		{
			name:  "replaces repeated placeholder",
			rules: map[string]interface{}{"word": "Go"},
			input: "%word% and %word%",
			want:  "Go and Go",
		},
		{
			name:  "returns input when rules are empty",
			rules: map[string]interface{}{},
			input: "Keep %name% as is",
			want:  "Keep %name% as is",
		},
		{
			name:  "leaves text without placeholders unchanged",
			rules: map[string]interface{}{"name": "Alice"},
			input: "Plain text",
			want:  "Plain text",
		},
		{
			name:  "preserves unmatched opening marker",
			rules: map[string]interface{}{"name": "Alice"},
			input: "Hello %name",
			want:  "Hello %name",
		},
		{
			name:  "replaces adjacent placeholders",
			rules: map[string]interface{}{"first": "A", "second": "B"},
			input: "%first%%second%",
			want:  "AB",
		},
	}

	f := NewEmptyFormatter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f.SetRules(tt.rules)
			got := f.Format(tt.input)
			if got != tt.want {
				t.Errorf("Format(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestFormatterFormatWithMarker(t *testing.T) {
	tests := []struct {
		name  string
		rules map[string]interface{}
		marker string
		input string
		want  string
	}{
		{
			name:  "replaces string value",
			rules: map[string]interface{}{"name": "Alice"},
			marker: "|",
			input: "Hello, |name|!",
			want:  "Hello, Alice!",
		},
		{
			name:  "replaces numeric value",
			rules: map[string]interface{}{"number": 123},
			marker: "1",
			input: "Number: 1number1",
			want:  "Number: 123",
		},
		{
			name: "calls function value",
			rules: map[string]interface{}{
				"func": func() string { return "hello world" },
			},
			marker: "%",
			input: "Result: %func%",
			want:  "Result: hello world",
		},
		{
			name:  "keeps unknown placeholder",
			rules: map[string]interface{}{"text": "12:00:00"},
			marker: "/",
			input: "Date: /date/",
			want:  "Date: /date/",
		},
		{
			name: "replaces known placeholders and keeps unknown ones",
			rules: map[string]interface{}{
				"text":   "12:00:00",
				"number": 123,
				"func":   func() string { return "hello world" },
			},
			marker: ".",
			input: "Text: .text. | Num: .number. | Func: .func. | Date: .date.",
			want:  "Text: 12:00:00 | Num: 123 | Func: hello world | Date: .date.",
		},
		{
			name:  "replaces repeated placeholder",
			rules: map[string]interface{}{"word": "Go"},
			marker: "x",
			input: "xwordx and xwordx",
			want:  "Go and Go",
		},
		{
			name:  "replaces adjacent placeholders",
			rules: map[string]interface{}{"first": "A", "second": "B"},
			marker: "|",
			input: "|first||second|",
			want:  "AB",
		},
	}

	f := NewEmptyFormatter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f.SetRules(tt.rules)
			f.SetMarkers(tt.marker, tt.marker)
			got := f.Format(tt.input)
			if got != tt.want {
				t.Errorf("Format(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}