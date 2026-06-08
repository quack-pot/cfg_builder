package internal

import (
	"encoding/json"
	"strings"
)

func IsLiteralJSON(value string) bool {
	normalized := strings.TrimSpace(value)

	if normalized == "true" || normalized == "false" || normalized == "null" {
		return true
	}

	if strings.HasPrefix(normalized, "[") && strings.HasSuffix(normalized, "]") {
		return true
	}

	var n json.Number
	return json.Unmarshal([]byte(normalized), &n) == nil
}
