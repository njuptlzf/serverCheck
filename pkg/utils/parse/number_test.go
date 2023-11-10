package parse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDiskForDir(t *testing.T) {
	tests := []struct {
		name      string
		element   string
		dir       string
		expSize   int64
		failedSug string
		failed    bool
	}{
		{
			name:      "valid input",
			element:   "/var/log;1G; >= 1G",
			dir:       "/var/log",
			expSize:   1073741824,
			failedSug: " >= 1G",
			failed:    false,
		},
		{
			name:      "valid input",
			element:   "/var/log;1GiB; >= 1G",
			dir:       "/var/log",
			expSize:   1073741824,
			failedSug: " >= 1G",
			failed:    false,
		},
		{
			name:      "valid input",
			element:   "/var/log;1GB; >= 1G",
			dir:       "/var/log",
			expSize:   1073741824,
			failedSug: " >= 1G",
			failed:    false,
		},
		{
			name:      "valid input",
			element:   "/var/log;1g; >= 1G",
			dir:       "/var/log",
			expSize:   1073741824,
			failedSug: " >= 1G",
			failed:    false,
		},
		{
			name:      "valid input",
			element:   "/var/log;1gB; >= 1G",
			dir:       "/var/log",
			expSize:   1073741824,
			failedSug: " >= 1G",
			failed:    false,
		},
		{
			name:    "invalid input",
			element: "invalid_input",
			dir:     "",
			expSize: 0,
			failed:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, expSize, failedSug, err := ParseDiskForDir(tt.element)
			assert.Equal(t, tt.dir, dir)
			assert.Equal(t, tt.expSize, expSize)
			assert.Equal(t, tt.failedSug, failedSug)
			if tt.failed {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseSize(t *testing.T) {
	tests := []struct {
		name     string
		size     float64
		expected string
	}{
		{
			name:     "bytes",
			size:     1024,
			expected: "1KiB",
		},
		{
			name:     "kilobytes",
			size:     1.1 * 1024 * 1024,
			expected: "1.1MiB",
		},
		{
			name:     "megabytes",
			size:     1.4 * 1024 * 1024 * 1024,
			expected: "1.4GiB",
		},
		{
			name:     "gigabytes",
			size:     1.6 * 1024 * 1024 * 1024 * 1024,
			expected: "1.6TiB",
		},
		{
			name:     "terabytes",
			size:     1024 * 1024 * 1024 * 1024 * 1024,
			expected: "1PiB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseSize(tt.size)
			assert.Equal(t, tt.expected, result)
		})
	}
}
