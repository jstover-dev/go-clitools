package common

import (
	"bufio"
	"os"
)

type FileReader struct {
	filePath string
	fileHandle *os.File
	reader *bufio.Reader
}

func NewFileReader(filepath string) (*FileReader, error) {
	var err error
	var file *os.File

	if filepath == "-" {
		file = os.Stdin
	} else {
		file, err = os.Open(filepath)
		if err != nil {
			return nil, err
		}
	}
	fileReader := FileReader{
		filepath,
		file,
		bufio.NewReader(file),
	}
	return &fileReader, nil
}

func (f FileReader) Close() error {
	return f.fileHandle.Close()
}

func (f FileReader) ReadByte() (byte, error) {
	return f.reader.ReadByte()
}