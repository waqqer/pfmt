package pfmt

import (
	"fmt"
	"time"
)

func DateRule() map[string]interface{} {
	return map[string]interface{}{
		"date": func() string {
			return time.Now().Format("2006-01-02")
		},
	}
}

func TimeRule() map[string]interface{} {
	return map[string]interface{}{
		"time": func() string {
			return time.Now().Format("15:04:05")
		},
		"time12": func() string {
			return time.Now().Format("03:04:05 PM")
		},
	}
}

func UnixRule() map[string]interface{} {
	return map[string]interface{}{
		"unix": func() string {
			return fmt.Sprintf("%d", time.Now().Unix())
		},
		"nano": func() string {
			return fmt.Sprintf("%d", time.Now().UnixNano())
		},
	}
}
