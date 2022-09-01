package main

import (
	"bytes"
	_ "embed"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	env = Environment{
		"BAR":   EnvValue{Value: "bar"},
		"EMPTY": EnvValue{Value: ""},
		"FOO": EnvValue{Value: `   foo
with new line`},
		"UNSET": EnvValue{Value: "", NeedRemove: true},
		"HELLO": EnvValue{Value: "\"hello\""},
	}

	out = `HELLO is ("hello")
BAR is (bar)
FOO is (   foo
with new line)
UNSET is ()
ADDED is ()
EMPTY is ()
arguments are arg1 arg2 arg3
`
)

func TestRunCmd(t *testing.T) {
	t.Run("Successful execution with ret-code 0", func(t *testing.T) {
		stdOut := os.Stdout

		reader, writer, err := os.Pipe()
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout = writer

		retCode := RunCmd([]string{"testdata/echo.sh", "arg1", "arg2", "arg3"}, env)
		require.Equal(t, 0, retCode)

		writer.Close()
		os.Stdout = stdOut

		var buffer bytes.Buffer
		_, err = io.Copy(&buffer, reader)
		if err != nil {
			log.Fatal(err)
		}

		require.Equal(t, out, buffer.String())
	})

	t.Run("return code no zero", func(t *testing.T) {
		returnCode := RunCmd([]string{"zero", "testdata/echo.sh"}, nil)
		require.Equal(t, returnCode, 0)
	})
}
