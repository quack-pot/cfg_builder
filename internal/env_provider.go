package internal

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type t_ConfigProviderENV struct {
	filenames []string
}

func NewConfigProviderENV(filenames ...string) IConfigProvider {
	return &t_ConfigProviderENV{
		filenames: filenames,
	}
}

func (p *t_ConfigProviderENV) Load() (map[string]any, error) {
	if err := godotenv.Load(p.filenames...); err != nil {
		return nil, err
	}

	var data map[string]any = make(map[string]any)
	for _, raw_env := range os.Environ() {
		split_env := strings.Split(raw_env, "=")

		if len(split_env) < 2 {
			continue
		}

		key := strings.ToLower(strings.TrimSpace(split_env[0]))
		value := strings.TrimSpace(split_env[1])

		sub_keys := strings.Split(key, string(CONFIG_KEY_SPLITTER))
		last_key_index := len(sub_keys) - 1

		var value_map map[string]any = data
		for idx := range last_key_index {
			sub_key := sub_keys[idx]

			sub_value, has_sub_map := value_map[sub_key]
			if has_sub_map {
				sub_map, is_sub_map := sub_value.(map[string]any)

				if is_sub_map {
					value_map = sub_map
					continue
				}
			}

			value_map[sub_key] = make(map[string]any)
			value_map = value_map[sub_key].(map[string]any)
		}

		value_map[sub_keys[last_key_index]] = value
	}

	return data, nil
}
