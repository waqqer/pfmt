package main

type placeholder_rule interface {
	func() string | string
}

type Formatter[T placeholder_rule] struct {
	rules map[string]T
}

func NewFormatter[T placeholder_rule](rules map[string]T) *Formatter[T] {
	return &Formatter[T]{
		rules: rules,
	}
}

func (f *Formatter[T]) Format(text string) string {
	if len(f.rules) == 0 {
		return text
	}
	return ""
}