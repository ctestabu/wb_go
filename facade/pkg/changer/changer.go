package changer

import (
	"errors"
	"fmt"
	"strings"

	v1 "github.com/wb_go/facade/pkg/api/v1"
)

// Changer where file extension changes
type Changer interface {
	ChangeFileName(fileName, format string) (msg string, err error)
	GetFileToChange() (str string)
}

type changer struct {
	filesToChange string
}

// ChangeFileName gets filename and changes its format
func (c *changer) ChangeFileName(fileName, format string) (msg string, err error) {
	msg = fmt.Sprintf(v1.FileToConvert, fileName)
	if fileName == "" {
		err = errors.New(v1.InvalidParameters)
		return
	}
	if strings.HasSuffix(fileName, ".old") {
		fileName = strings.Replace(fileName, ".old", format, 1)
	} else {
		err = errors.New(v1.WrongSuffix)
		return
	}
	c.filesToChange = fileName
	return
}

// GetFileToChange returns modified in ChangeFileName fileName
func (c *changer) GetFileToChange() (str string) {
	return c.filesToChange
}

//NewChanger changer entity
func NewChanger(fileName string) Changer {
	return &changer{
		filesToChange: fileName,
	}
}
