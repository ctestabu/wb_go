package changer

import (
	"errors"
	v1 "facade/pkg/api/v1"
	"fmt"
	"strings"
)

//Methods for Changer interface
type Changer interface {
	ChangeFileName(fileName, format string) (msg string, err error)
	GetFileToChange()(str string)
}

type changer struct {
	filesToChange string
}

//ChangeFileName gets filename and changes its format
func (c *changer) ChangeFileName(fileName, format string)(msg string, err error) {
	msg = fmt.Sprintf(v1.FileToConvert, fileName)
	if fileName == "" {
		err = errors.New(v1.InvalidParameters)
	}
	if strings.HasSuffix(fileName, ".old") {
		fileName = strings.Replace(fileName, ".old", format, 1)
	}
	c.filesToChange = fileName
	return
}

//GetFileToChange returns modified in ChangeFileName fileName
func (c *changer) GetFileToChange ()(str string) {
	return c.filesToChange
}

//Creates new changer entity
func New(fileName string) Changer {
	return &changer{
		filesToChange: fileName,
	}
}
