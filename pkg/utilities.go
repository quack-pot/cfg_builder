package cfg_builder

import (
	"regexp"
	"strings"

	"github.com/quack-pot/cfg_builder/internal"
)

func BuildKey(parts ...string) string {
	if len(parts) == 0 {
		return ""
	}

	key := ""
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		key += part + string(internal.CONFIG_KEY_SPLITTER)
	}

	remove_duplicate_regex := regexp.MustCompile(
		regexp.QuoteMeta(string(internal.CONFIG_KEY_SPLITTER)) + `{2,}`,
	)

	key = remove_duplicate_regex.ReplaceAllString(key, string(internal.CONFIG_KEY_SPLITTER))

	key = strings.ToLower(strings.Trim(key, string(internal.CONFIG_KEY_SPLITTER)))

	return key
}
