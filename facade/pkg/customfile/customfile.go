package customfile

import (
	"os"
)

// CustomFile creates file
type CustomFile interface {
	GiveConvertedFile(fileName string) (err error)
}

type customFile struct {
	fileNames string
}

// Creates file with modified format
func (c *customFile) GiveConvertedFile(fileName string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		return
	}

	err = file.Close()
	if err != nil {
		return
	}
	return
}

// NewCustomFile creates custom file entity
func NewCustomFile(fileNames string) CustomFile {
	return &customFile{
		fileNames: fileNames,
	}
}
