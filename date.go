package pananames

import (
	"strings"
	"time"
)

// Custom datetime structures need to avoid JSON unmarshall errors when DateTime in JSON returns as empty string ""

type PnTime struct {
	time.Time
}

// Custom Unmarshall for PnTime
func (t *PnTime) UnmarshalJSON(data []byte) error {
	if string(data) == `""` {
		return nil
	}
	return t.Time.UnmarshalJSON(data)
}

type PnDate struct {
	time.Time
}

// Custom Unmarshall for PnDate
func (d *PnDate) UnmarshalJSON(data []byte) error {
	val := strings.Trim(string(data), `"`)
	if val == "" {
		return nil
	}

	var err error
	if d.Time, err = time.Parse(time.DateOnly, val); err != nil {
		return err
	}

	return nil
}
