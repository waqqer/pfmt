package main

type placeholder_rule interface {
	func() string | string
}

type Formatter[T placeholder_rule] struct {
	rules  map[string]T
	marker string
}

func NewFormatter[T placeholder_rule](rules map[string]T) *Formatter[T] {
	return &Formatter[T]{
		rules:  rules,
		marker: "%",
	}
}

func NewMarkerFormatter[T placeholder_rule](rules map[string]T, marker string) *Formatter[T] {
	return &Formatter[T]{
		rules:  rules,
		marker: marker,
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
