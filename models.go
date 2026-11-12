package pfmt

type Rule map[string]interface{}

type Pfmt struct {
	rules        Rule
	start_marker string
	end_marker   string
}

func NewEmptyFormatter() *Pfmt {
	return &Pfmt{
		rules:        Rule{},
		start_marker: "%",
		end_marker:   "%",
	}
}

func NewFormatter(rules ...Rule) *Pfmt {
	return &Pfmt{
		rules:        margeRules(rules...),
		start_marker: "%",
		end_marker:   "%",
	}
}
