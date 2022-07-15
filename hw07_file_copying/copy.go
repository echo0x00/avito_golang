package main

import (
	"errors"
	"io"
	"math"
	"os"

	"github.com/echo0x00/avito_golang/hw07_file_copying/internal/bar"
)

var (
	ErrFileCreate            = errors.New("file create error")
	ErrUnsupportedFile       = errors.New("unsupported file or not exists")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath *string, offset, limit *int64) error {
	sourceFileInfo, err := os.Stat(*fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	sourceFileSize := sourceFileInfo.Size()
	if sourceFileSize < *offset {
		return ErrOffsetExceedsFileSize
	}

	if sourceFileSize == 0 || !sourceFileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	sourceFile, _ := os.OpenFile(*fromPath, os.O_RDONLY, 0o644)
	defer sourceFile.Close()

	if *offset > 0 {
		_, _ = sourceFile.Seek(*offset, io.SeekStart)
	}

	toFile, err := os.Create(*toPath)
	if err != nil {
		return ErrFileCreate
	}
	defer toFile.Close()

	if *limit == 0 || *limit > sourceFileSize {
		*limit = sourceFileSize
	}

	bufferSize := int64(math.Ceil(float64(*limit) / 10))
	maxBufferSize := int64(1 << 20)
	if bufferSize > maxBufferSize {
		bufferSize = maxBufferSize
	}

	bar := bar.NewProgress(limit)

	for totalWritten := int64(0); totalWritten < *limit; {
		written, err := io.CopyN(toFile, sourceFile, bufferSize)
		if err != nil {
			if errors.Is(err, io.EOF) {
				_ = bar.Finish()
				break
			}

			return err
		}

		totalWritten += written
		_ = bar.Step(totalWritten)
	}

	return nil
}
