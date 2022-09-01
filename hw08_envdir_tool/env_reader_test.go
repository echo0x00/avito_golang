package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("Test ReadDir", func(t *testing.T) {
		expected := Environment{
			"BAR":   EnvValue{Value: "bar"},
			"EMPTY": EnvValue{Value: ""},
			"FOO": EnvValue{Value: `   foo
with new line`},
			"UNSET": EnvValue{Value: "", NeedRemove: true},
			"HELLO": EnvValue{Value: "\"hello\""},
		}

		env, err := ReadDir("./testdata/env")
		require.NoError(t, err)
		require.Equal(t, expected, env)
	})

	t.Run("Test Dir Not Found", func(t *testing.T) {
		_, err := ReadDir("./not_found!")
		require.ErrorAs(t, err, &ErrDirNotFound)
	})
}
