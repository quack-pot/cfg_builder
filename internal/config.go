package internal

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"
)

type IConfig interface {
	Get(key string) (any, bool)
	Set(key string, value any)

	Bind(dst any, prefix string) error
}

type t_Config struct {
	data map[string]any
}

var ErrInvalidDst error = errors.New("Destination must be pointer type and non-nil.")

func NewConfig(data map[string]any) IConfig {
	return &t_Config{
		data: data,
	}
}

func (c *t_Config) Get(key string) (any, bool) {
	value_map := c.data
	split_keys := strings.Split(key, string(CONFIG_KEY_SPLITTER))
	last_key_index := len(split_keys) - 1

	for idx := range last_key_index {
		sub_key := split_keys[idx]
		sub_value, has_value := value_map[sub_key]

		if !has_value {
			return nil, false
		}

		sub_map, is_sub_map := sub_value.(map[string]any)

		if !is_sub_map {
			return nil, false
		}

		value_map = sub_map
	}

	value, has_value := value_map[split_keys[last_key_index]]
	return value, has_value
}

func (c *t_Config) Set(key string, value any) {
	value_map := c.data
	split_keys := strings.Split(key, string(CONFIG_KEY_SPLITTER))
	last_key_index := len(split_keys) - 1

	for idx := range last_key_index {
		sub_key := split_keys[idx]
		sub_value, has_value := value_map[sub_key]

		if !has_value {
			value_map[sub_key] = make(map[string]any)
			value_map = value_map[sub_key].(map[string]any)
			continue
		}

		sub_map, is_sub_map := sub_value.(map[string]any)

		if is_sub_map {
			value_map = sub_map
			continue
		}

		value_map[sub_key] = make(map[string]any)
		value_map = value_map[sub_key].(map[string]any)
	}

	value_map[split_keys[last_key_index]] = value
}

func (c *t_Config) Bind(dst any, prefix string) error {
	dst_val := reflect.ValueOf(dst)

	if dst_val.IsNil() || dst_val.Kind() != reflect.Ptr {
		return ErrInvalidDst
	}

	value_map := c.data

	if prefix != "" {
		split_keys := strings.Split(prefix, string(CONFIG_KEY_SPLITTER))
		for idx := range len(split_keys) {
			sub_key := split_keys[idx]
			sub_value, has_value := value_map[sub_key]

			if !has_value {
				return nil
			}

			sub_map, is_sub_map := sub_value.(map[string]any)

			if is_sub_map {
				value_map = sub_map
				continue
			}

			return nil
		}
	}

	json_intermediate, err := json.Marshal(value_map)
	if err != nil {
		return err
	}

	return json.Unmarshal(json_intermediate, dst)
}
