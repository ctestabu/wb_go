package customFile

import (
	"errors"
	"os"
)

type CustomFile interface {
	//giveConvertedFile
	GiveConvertedFile(fileName string) (err error)
}

type customFile struct {
	fileNames []string
}

func (c *customFile) GiveConvertedFile(fileName string) (err error) {
	file, err := os.Create(fileName)
	if err != nil {
		err =  errors.New("Too bad")
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