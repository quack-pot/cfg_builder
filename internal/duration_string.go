package internal

import (
	"encoding/json"
	"time"
)

type DurationString time.Duration

func (d *DurationString) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	duration, err := time.ParseDuration(raw)
	if err != nil {
		return err
	}

	*d = DurationString(duration)
	return nil
}
