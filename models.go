package pfmt

type Pfmt struct {
	rules  map[string]interface{}
	start_marker string
	end_marker string
}

func NewEmptyFormatter() *Pfmt {
	return &Pfmt{
		rules: map[string]interface{}{},
		start_marker: "%",
		end_marker: "%",
	}
}

func NewFormatter(rules ...map[string]interface{}) *Pfmt {
	return &Pfmt{
		rules: margeMaps(rules...),
		start_marker: "%",
		end_marker: "%",
	}
}