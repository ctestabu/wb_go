package customstrslice

import (
	"errors"
	v1 "github.com/wb_go/visitor/pkg/api/v1"
	"strings"
)

type visitor interface {
	ReplaceStrSlice(c CustomStrSlice) (err error)
}

// CustomStrSlice describes custom slice methods
type CustomStrSlice interface {
	ToUpper()
	ToLower()
	LenSlice() (l int)
	GetFirstElem() (cs string)
	ReplaceFirstElem(str string) (err error)
	Accept(v visitor) (err error)
}

type customStrSlice struct {
	strSlice []string
}

// Accept allows visitor to function
func (c *customStrSlice) Accept(v visitor) (err error) {
	return v.ReplaceStrSlice(c)
}

// ToUpper ...
func (c *customStrSlice) ToUpper() {
	for i := range c.strSlice {
		c.strSlice[i] = strings.ToUpper(c.strSlice[i])
	}
}

// ToLower ...
func (c *customStrSlice) ToLower() {
	for i := range c.strSlice {
		c.strSlice[i] = strings.ToLower(c.strSlice[i])
	}
}

// GetFirstElem ...
func (c *customStrSlice) GetFirstElem() (cs string) {
	return c.strSlice[0]
}

// ReplaceFirstElem ...
func (c *customStrSlice) ReplaceFirstElem(str string) (err error) {

	if str == "" {
		err = errors.New(v1.ErrorEmptyReplace)
	}
	c.strSlice[0] = str
	return
}

// LenSlice calls standard len to return len of custom slice
func (c *customStrSlice) LenSlice() (l int) {
	return len(c.strSlice)
}

// NewCustomStrSlice creates new entity of custom slice
func NewCustomStrSlice(slice []string) CustomStrSlice {
	return &customStrSlice{
		strSlice: slice,
	}
}
