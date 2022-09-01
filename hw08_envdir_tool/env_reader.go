package main

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

var (
	ErrDirNotFound = errors.New("dir is not found")
	ErrFileInfo    = errors.New("get file info error")
)

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, ErrDirNotFound
	}

	env := make(Environment, len(files))

	for _, fileEntry := range files {
		var envVal EnvValue

		filePath := filepath.Join(dir, fileEntry.Name())
		file, _ := os.Open(filePath)

		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return env, ErrFileInfo
		}
		if size := fileInfo.Size(); size == 0 {
			envVal.NeedRemove = true
			env[fileEntry.Name()] = envVal
			return env, nil
		}

		if err != nil {
			os.Stderr.WriteString(err.Error())
		} else {
			scanner := bufio.NewReader(file)
			if line, _, err := scanner.ReadLine(); err == nil {
				lineString := bytes.ReplaceAll(line, []byte{0}, []byte("\n"))
				value := strings.Trim(string(lineString), "\t")
				if value != " " {
					envVal.Value = value
				}
			}
		}

		env[fileEntry.Name()] = envVal
	}

	return env, err
}
