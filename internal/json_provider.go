package internal

import (
	"encoding/json"
	"os"
)

type t_ConfigProviderJSON struct {
	filenames []string
}

func NewConfigProviderJSON(filenames ...string) IConfigProvider {
	return &t_ConfigProviderJSON{
		filenames: filenames,
	}
}

func (p *t_ConfigProviderJSON) Load() (map[string]any, error) {
	var data map[string]any = make(map[string]any)

	for _, filename := range p.filenames {
		file, err := os.Open(filename)

		if err != nil {
			return nil, err
		}

		decoder := json.NewDecoder(file)

		if err := decoder.Decode(&data); err != nil {
			return nil, err
		}

		if err := file.Close(); err != nil {
			return nil, err
		}
	}

	return data, nil
}
