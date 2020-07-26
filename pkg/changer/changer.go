package changer

import (
	"errors"
	v1 "facade/pkg/api/v1"
	"fmt"
	"strings"
)

type Changer interface {
	ChangeFileName(fileName, format string) (msg string, err error)
	ListFilesToChange() (msg string, err error)
}

type changer struct {
	filesToChange []string
}

func (c *changer) ChangeFileName(fileName, format string)(msg string, err error) {
	if fileName == "" /*|| (format != ".custom1" && format != ".custom2")*/ {
		err = errors.New(v1.InvalidParameters)
	}
	if strings.HasSuffix(fileName, ".old") {
		fileName = strings.Replace(fileName, ".old", format, 1)
	}
	c.filesToChange = append(c.filesToChange, fileName)
	msg = fmt.Sprintf(v1.FileToConvert, fileName)
	return
}

func (c *changer) ListFilesToChange() (msg string, err error) {
	if len(c.filesToChange) == 0 {
		err = errors.New(v1.FilesToChangeError)
	}
	msg = fmt.Sprintf(v1.FileListOut, strings.Join(c.filesToChange, v1.FilesSeparator))
	return
}

func New(fileNames []string) Changer {
	return &changer{
		filesToChange: fileNames,
	}
}