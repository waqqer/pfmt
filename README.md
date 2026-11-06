# pfmt

`pfmt` is a lightweight Go library for formatting text with configurable placeholders. It replaces placeholders whose keys are present in the formatter's rules and leaves unknown placeholders unchanged. Rule values can be strings, other values formatted with `%v`, or functions returning `string` or `interface{}`.

## Installation

```sh
go get github.com/waqqer/pfmt/v2
```

## Examples

### Replace values

```go
package main

import (
	"fmt"
	
	"github.com/waqqer/pfmt/v2"
)

func main() {
	rules := map[string]interface{}{
		"name":  "Alice",
		"items": 3,
	}

	formatter := pfmt.NewFormatter(rules)
	text := formatter.Format("Hello, %name%! You have %items% items.")

	fmt.Println(text)
	// Hello, Alice! You have 3 items.
}
```

### Merge and update rules

`NewFormatter` and `AddRule` accept one or more rule maps. When a key appears more than once, the value from the later map replaces the earlier value. `SetRules` replaces all current rules, while `Reset` removes them.

```go
formatter := pfmt.NewFormatter(
	map[string]interface{}{"name": "Alice", "city": "London"},
	map[string]interface{}{"name": "Bob"},
)
formatter.AddRule(map[string]interface{}{"items": 3})

text := formatter.Format("%name% lives in %city% and has %items% items.")
// Bob lives in London and has 3 items.
```

### Use a function and keep unknown placeholders

```go
package main

import (
	"fmt"
	"time"

	"github.com/waqqer/pfmt/v2"
)

func main() {
	rules := map[string]interface{}{
		"time": func() string {
			return time.Now().Format("15:04:05")
		},
	}

	formatter := pfmt.NewFormatter(rules)
	text := formatter.Format("Current time: %time%. Date: %date%")

	fmt.Println(text)
	// Current time: 15:04:05. Date: %date%
}
```

Functions returning `interface{}` are also supported and their result is formatted with `%v`.

### Empty formatter

```go
formatter := pfmt.NewEmptyFormatter()
text := formatter.Format("Current time: %time%")
// Current time: %time%
```

### Change formatter markers

```go
formatter := pfmt.NewFormatter(map[string]interface{}{"name": "Alice"})
formatter.SetMarkers("{", "}")
text := formatter.Format("Hello, {name}!")
// Hello, Alice!
```
