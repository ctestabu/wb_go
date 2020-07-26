package facade

import (
	v1 "facade/pkg/api/v1"
	"strings"
)

type changer interface {
	ChangeFile(fileName, format string) (msg string, err error) // change file format
	ListFilesToChange() (msg string, err error) // Lst of changed files
}

type customFile interface {
	GiveConvertedFile(fileName string) (err error) // Creates file
}

//Control interface to change file extension
type Converter interface {
	Receive(fileName string, format string) (msg string, err error)
}

type converter struct {
	files    customFile
	toChange changer
}

func (c *converter) Receive(fileName string, format string) (msg string, err error) {
	var logger = make([]string, 0)
	var log string

	log,err = c.toChange.ChangeFile(fileName, format)
	if err != nil {
		return
	}
	logger = append(logger, log)

	log,err = c.toChange.ListFilesToChange()
	if err != nil {
		return
	}
	logger = append(logger, log)

	err = c.files.GiveConvertedFile(fileName)
	if err != nil {
		return
	}
	msg = strings.Join(logger, v1.LogSeparator)
	return
}

//Creates new Converter entity
func NewConverter(files customFile, toChange changer) Converter {
	return &converter{
		files:    files,
		toChange: toChange,
	}
}
