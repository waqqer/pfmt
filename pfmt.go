package pfmt

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
	return "sd"
}