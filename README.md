# pfmt

`pfmt` is a lightweight Go library for formatting text with `%key%` placeholders. It replaces placeholders whose keys are present in the supplied rules and leaves unknown placeholders unchanged. String-returning functions can also be used as values.

## Installation

```sh
go get github.com/waqqer/pfmt
```

## Examples

### Replace values

```go
package main

import (
	"fmt"

	"github.com/waqqer/pfmt"
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

### Use a function and keep unknown placeholders

```go
package main

import (
	"fmt"
	"time"

	"github.com/waqqer/pfmt"
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
