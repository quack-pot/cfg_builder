package internal

import (
	"encoding/json"
	"fmt"
	"log"
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

func EnvDataToJSON(data map[string]any) (string, error) {
	var inner_content []string = make([]string, len(data))

	var idx int = 0
	for key := range data {
		inner_data := data[key]

		switch inner_data := inner_data.(type) {
		case map[string]any:
			inner_map_content, err := EnvDataToJSON(inner_data)

			if err != nil {
				return "", err
			}

			inner_content[idx] = fmt.Sprintf("%q:%s", key, inner_map_content)

		case string:
			if IsLiteralJSON(inner_data) {
				inner_content[idx] = fmt.Sprintf("%q:%s", key, inner_data)
				break
			}

			inner_content[idx] = fmt.Sprintf("%q:%q", key, inner_data)

		default:
			log.Panicf("[Error]: Unexpected ENV type (%v)", inner_data)
		}

		idx++
	}

	return fmt.Sprintf("{%s}", strings.Join(inner_content, ",")), nil
}

func (p *t_ConfigProviderENV) loadEnvironment(
	environ map[string]string,
) (map[string]any, error) {
	var data map[string]any = make(map[string]any)
	for raw_key, raw_value := range environ {
		key := strings.TrimSpace(raw_key)
		value := strings.TrimSpace(raw_value)

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

	json_raw, err := EnvDataToJSON(data)
	if err != nil {
		return nil, err
	}

	var result map[string]any = make(map[string]any)
	if err := json.Unmarshal([]byte(json_raw), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (p *t_ConfigProviderENV) Load() (map[string]any, error) {
	if len(p.filenames) > 0 {
		environ, err := godotenv.Read(p.filenames...)

		if err != nil {
			return nil, err
		}

		return p.loadEnvironment(environ)
	}

	var environ map[string]string = make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)

		if len(parts) < 2 {
			log.Printf("[Warning]: Malformed environment variable being skipped (%s).", env)
			continue
		}

		environ[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	return p.loadEnvironment(environ)
}
