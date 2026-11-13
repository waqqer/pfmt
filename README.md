# pfmt

`pfmt` is a lightweight Go library for formatting text with configurable placeholders. It replaces placeholders whose keys are present in a rule set and leaves unknown placeholders unchanged. Values are formatted with `fmt`, and rule values can be functions returning `string` or `interface{}`.

## Installation

```sh
go get github.com/waqqer/pfmt/v2
```

## Formatting

Use the package-level `Format` function to format text with the default `%` markers:

```go
package main

import (
	"fmt"

	"github.com/waqqer/pfmt/v2"
)

func main() {
	rules := pfmt.Rule{
		"name":  "Alice",
		"items": 3,
	}

	text := pfmt.Format("Hello, %name%! You have %items% items.", rules)
	fmt.Println(text)
	// Hello, Alice! You have 3 items.
}
```

## Formatter

Create a `Pfmt` formatter with `NewFormatter`, or use `NewEmptyFormatter` when you want to add rules later. `NewFormatter` accepts multiple rule sets. If a key appears more than once, the value from the later set takes precedence.

```go
formatter := pfmt.NewFormatter(
	pfmt.Rule{"name": "Alice", "city": "London"},
	pfmt.Rule{"name": "Bob"},
)

text := formatter.Format("%name% lives in %city%.")
// Bob lives in London.
```

### Add, replace, and reset rules

`AddRule` merges one or more rule sets into the formatter. Later values override earlier ones. `SetRules` replaces all current rules, and `Reset` removes all rules.

```go
formatter.AddRule(pfmt.Rule{"items": 3})
formatter.SetRules(pfmt.Rule{"name": "Carol"})
formatter.Reset()
```

### Custom markers

Use `SetMarkers` to choose different start and end markers. Empty or whitespace-only markers are ignored.

```go
formatter := pfmt.NewFormatter(pfmt.Rule{"name": "Alice"})
formatter.SetMarkers("{", "}")

text := formatter.Format("Hello, {name}!")
// Hello, Alice!
```

### Write formatted text

`WFormat` writes the formatted result to an `io.Writer`. `PFormat` assigns the result to a string pointer.

```go
var output strings.Builder
formatter.WFormat(&output, "Hello, {name}!")

text := ""
formatter.PFormat(&text, "Hello, {name}!")
```

## Function values

Functions are called when their placeholders are formatted. Functions returning `interface{}` are formatted with `%v`.

```go
rules := pfmt.Rule{
	"greeting": func() string { return "Hello" },
	"count":    func() interface{} { return 3 },
}

text := pfmt.Format("%greeting%! Items: %count%", rules)
// Hello! Items: 3
```

## Built-in time rules

The package provides `DateRule`, `TimeRule`, and `UnixRule`. Each returns a map of rules that can be converted to `pfmt.Rule` and merged with other rule sets.

```go
formatter := pfmt.NewFormatter(pfmt.Rule(pfmt.DateRule()))
formatter.AddRule(pfmt.Rule(pfmt.TimeRule()))

text := formatter.Format("Date: %date%, time: %time%")
```

The helpers provide these keys:

- `DateRule`: `date` (`YYYY-MM-DD`)
- `TimeRule`: `time` (`HH:MM:SS`) and `time12` (12-hour time with AM/PM)
- `UnixRule`: `unix` (Unix timestamp in seconds) and `nano` (Unix timestamp in nanoseconds)
