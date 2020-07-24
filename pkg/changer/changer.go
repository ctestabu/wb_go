package changer

import (
	"errors"
	"fmt"
	"strings"
)

type Changer interface {
	ChangeFile(fileName, format string) (msg string, err error)
	ListFilesToChange() (msg string, err error)
}

type changer struct {
	filesToChange []string
}

func (c *changer) ChangeFile(fileName, format string)(msg string, err error) {
	if fileName == "" || format != ".custom1" || format != ".custom" {
		err = errors.New("Invalid parameters")
	}
	if strings.HasSuffix(fileName, ".old1") || strings.HasSuffix(fileName, ".old2") {
		strings.Replace(fileName, ".old1", ".custom1", 1)
		strings.Replace(fileName, ".old2", ".custom2", 1)
	}
	c.filesToChange = append(c.filesToChange, fileName)
	msg = fmt.Sprint("File converted")
	return
}

func (c *changer) ListFilesToChange() (msg string, err error) {
	if len(c.filesToChange) == 0 {
		err = errors.New("Too bad")
	}
	msg = fmt.Sprint("list of files")
	return
}

func New(fileNames []string) Changer {
	return &changer{
		filesToChange: fileNames,
	}
}