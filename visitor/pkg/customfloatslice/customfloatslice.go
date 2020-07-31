package customfloatslice

import (
	"errors"
	"math"
)

type visitor interface {
	RoundAndReplace(c CustomFloatSlice) (err error)
}

// CustomFloatSlice describes custom slice methods
type CustomFloatSlice interface {
	Round() (err error)
	SwapFirstAndLast()
	GetFirstElem() (el float64)
	ReplaceFirstElem(el float64)
	LenSlice() (l int)
	Accept(v visitor) (err error)
}

type customFloatSlice struct {
	fltSlice []float64
}

// Accept allows visitor to function
func (c *customFloatSlice) Accept(v visitor) (err error) {
	return v.RoundAndReplace(c)
}

// Round rounds all elements in custom slice
func (c *customFloatSlice) Round() (err error) {
	if len(c.fltSlice) == 0 {
		err = errors.New("Dont even bother to do something here")
		return
	}
	for i, j := range c.fltSlice {
		c.fltSlice[i] = math.Round((j * 100) / 100)
	}
	return
}

// SwapFirstAndLast ...
func (c *customFloatSlice) SwapFirstAndLast() {
	c.fltSlice[0], c.fltSlice[len(c.fltSlice)-1] = c.fltSlice[len(c.fltSlice)-1], c.fltSlice[0]
}

// GetFirstElem ...
func (c *customFloatSlice) GetFirstElem() (el float64) {
	return c.fltSlice[0]
}

// LenSlice calls standard len to return len of custom slice
func (c *customFloatSlice) LenSlice() (l int) {
	return len(c.fltSlice)
}

// ReplaceFirstElem ...
func (c *customFloatSlice) ReplaceFirstElem(el float64) {
	c.fltSlice[0] = el
}

// NewCustomFloatSlice creates new entity of custom slice
func NewCustomFloatSlice(slice []float64) CustomFloatSlice {
	return &customFloatSlice{
		fltSlice: slice,
	}
}
