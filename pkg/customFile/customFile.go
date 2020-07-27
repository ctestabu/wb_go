package customFile

import (
	"errors"
	v1 "facade/pkg/api/v1"
	"os"
)

//Methods for CustomFile interface
type CustomFile interface {
	GiveConvertedFile(fileName string) (err error)
}

type customFile struct {
	fileNames string
}

//Creates file with modified format
func (c *customFile) GiveConvertedFile(fileName string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		err =  errors.New(v1.CreateFileError)
		os.Exit(1)
	}

	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()
	return
}

//Creates new custom file entity
func New(fileNames string) CustomFile {
	return &customFile{
		fileNames: fileNames,
	}
}
