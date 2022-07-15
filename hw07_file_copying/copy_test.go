package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	input, _ := os.CreateTemp(os.TempDir(), "input.*.txt")
	input.Write([]byte("test"))
	output, _ := os.CreateTemp(os.TempDir(), "output.*.txt")
	casesErrors := []struct {
		name     string
		fromPath string
		toPath   string
		limit    int64
		offset   int64
		expected error
	}{
		{
			name:     "File not exists",
			fromPath: "/home/fileNotExists",
			toPath:   output.Name(),
			limit:    0,
			offset:   0,
			expected: ErrUnsupportedFile,
		},
		{
			name:     "File exists but not support",
			fromPath: "/dev/random",
			toPath:   output.Name(),
			limit:    0,
			offset:   0,
			expected: ErrUnsupportedFile,
		},
		{
			name:     "Destination file path not support",
			fromPath: input.Name(),
			toPath:   ":~/",
			limit:    0,
			offset:   0,
			expected: ErrFileCreate,
		},
	}

	for _, c := range casesErrors {
		t.Run(c.name, func(t *testing.T) {
			err := Copy(&c.fromPath, &c.toPath, &c.offset, &c.limit)
			require.Truef(t, errors.Is(err, c.expected), "actual error %v", err)
		})
	}
}
