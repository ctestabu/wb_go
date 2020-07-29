package customstrslice

import "strings"

type visitor interface {
	ReplaceStrSlice(c CustomStrSlice) (err error)
}

type CustomStrSlice interface {
	ToUpper()
	ToLower()
	LenSlice() (l int)
	GetFirstElem() (cs string)
	ReplaceFirstElem(str string)
	Accept(v visitor) (err error)
}

type customStrSlice struct {
	strSlice []string
}

func (c *customStrSlice) Accept(v visitor) (err error) {
	return v.ReplaceStrSlice(c)
}

// TODO Methods here

func (c *customStrSlice) ToUpper() {
	for i, _ := range c.strSlice {
		c.strSlice[i] = strings.ToUpper(c.strSlice[i])
	}
}

func (c *customStrSlice) ToLower() {
	for i, _ := range c.strSlice {
		c.strSlice[i] = strings.ToLower(c.strSlice[i])
	}
}

func (c *customStrSlice) GetFirstElem() (cs string) {
	return c.strSlice[0]
}

func (c *customStrSlice) ReplaceFirstElem(str string) {
	c.strSlice[0] = str
}

func (c *customStrSlice) LenSlice() (l int) {
	return len(c.strSlice)
}

func NewCustomStrSlice(slice []string) CustomStrSlice {
	return &customStrSlice{
		strSlice: slice,
	}
}
