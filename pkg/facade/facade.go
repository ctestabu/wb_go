package facade

import (
	v1 "facade/pkg/api/v1"
	"strings"
)

type changer interface {
	ChangeFileName(fileName, format string) (msg string, err error)
	GetFileToChange()(str string)
}

type customFile interface {
	GiveConvertedFile(fileName string) (err error)
}

//Control interface to change file extension
type Converter interface {
	Receive(fileName string, format string) (msg string, err error)
}

type converter struct {
	files    customFile
	toChange changer
}
//Receive method gets string(name of file to be created), changes its extension to format
func (c *converter) Receive(fileName string, format string) (msg string, err error) {
	var logger = make([]string, 0)
	var log string

	log,err = c.toChange.ChangeFileName(fileName, format)
	if err != nil {
		return
	}
	logger = append(logger, log)

	err = c.files.GiveConvertedFile(c.toChange.GetFileToChange())
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
