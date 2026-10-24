package main

type placeholder_rule interface {
	func() string | string
}

type Formatter[T placeholder_rule] struct {
	rules  map[string]T
	start_marker string
	end_marker string
}

func NewFormatter[T placeholder_rule](rules map[string]T) *Formatter[T] {
	return &Formatter[T]{
		rules:  rules,
		start_marker: "%",
		end_marker: "%",
	}
}

func NewMarkerFormatter[T placeholder_rule](rules map[string]T, start_marker string, end_marker string) *Formatter[T] {
	return &Formatter[T]{
		rules:  rules,
		start_marker: start_marker,
		end_marker: end_marker,
	}
}

func (f *Formatter[T]) Format(text string) string {
	if len(f.rules) == 0 {
		return text
	}
	return f.replaceAll(text)
}

func (f *Formatter[T]) replaceAll(text string) string {
	return "sd"
}

func parseText(text string) []string {

}
