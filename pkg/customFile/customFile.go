package customFile

import (
	"errors"
	v1 "facade/pkg/api/v1"
	"fmt"
	"os"
)

type CustomFile interface {
	GiveConvertedFile(fileName string) (err error)
}

type customFile struct {
	fileNames []string
}

func (c *customFile) GiveConvertedFile(fileName string) (err error) {
	fmt.Println("-------------1",fileName)
	file, err := os.Create(fileName)
	if err != nil {
		err =  errors.New(v1.CreateFileError)
		os.Exit(1)
	}
	//  https://www.joeshaw.org/dont-defer-close-on-writable-files/
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()
	return
}

func New(fileNames []string) CustomFile {
	return &customFile{
		fileNames: fileNames,
	}
}
