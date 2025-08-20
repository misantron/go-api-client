package pananames

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPnTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		wantErr  string
		wantDate string
	}{
		{
			name:     "empty",
			input:    []byte(`""`),
			wantDate: "0001-01-01 00:00:00",
		},
		{
			name:    "invalid",
			input:   []byte(`"abc"`),
			wantErr: `parsing time "abc" as "2006-01-02T15:04:05Z07:00": cannot parse "abc" as "2006"`,
		},
		{
			name:     "valid",
			input:    []byte(`"2020-01-02T03:04:05Z"`),
			wantDate: "2020-01-02 03:04:05",
		},
	}

	var (
		s   *PnTime
		err error
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s = new(PnTime)
			err = s.UnmarshalJSON(tt.input)

			if err != nil {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.Equal(t, tt.wantDate, s.Time.Format(time.DateTime))
			}
		})
	}
}

func TestPnDate_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		wantErr  string
		wantDate string
	}{
		{
			name:     "empty",
			input:    []byte(`""`),
			wantDate: "0001-01-01",
		},
		{
			name:    "invalid_format",
			input:   []byte(`"02.03.2004"`),
			wantErr: `parsing time "02.03.2004" as "2006-01-02": cannot parse "02.03.2004" as "2006"`,
		},
		{
			name:    "invalid",
			input:   []byte(`"abc"`),
			wantErr: `parsing time "abc" as "2006-01-02": cannot parse "abc" as "2006"`,
		},
		{
			name:     "valid",
			input:    []byte(`"2020-01-02"`),
			wantDate: "2020-01-02",
		},
	}

	var (
		s   *PnDate
		err error
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s = new(PnDate)
			err = s.UnmarshalJSON(tt.input)

			if err != nil {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.Equal(t, tt.wantDate, s.Time.Format(time.DateOnly))
			}
		})
	}
}
