package customfloatslice

import "math"

type visitor interface {
	RoundAndSwap(c CustomFloatSlice) (err error)
}

type CustomFloatSlice interface {
	Round()
	SwapFirstAndLast()
	GetFirstElem() (el float64)
	ReplaceFirstElem(el float64)
	LenSlice() (l int)
	Accept(v visitor) (err error)
}

type customFloatSlice struct {
	fltSlice []float64
}

func (c *customFloatSlice) Accept(v visitor) (err error) {
	return v.RoundAndSwap(c)
}

func (c *customFloatSlice) Round() {
	for i, j := range c.fltSlice {
		c.fltSlice[i] = math.Round((j * 100) / 100)
	}
}

func (c *customFloatSlice) SwapFirstAndLast() {
	c.fltSlice[0], c.fltSlice[len(c.fltSlice)-1] = c.fltSlice[len(c.fltSlice)-1], c.fltSlice[0]
}

func (c *customFloatSlice) GetFirstElem() (el float64) {
	return c.fltSlice[0]
}

func (c *customFloatSlice) LenSlice() (l int) {
	return len(c.fltSlice)
}

func (c *customFloatSlice) ReplaceFirstElem(el float64) {
	c.fltSlice[0] = el
}

func NewCustomFloatSlice(slice []float64) CustomFloatSlice {
	return &customFloatSlice{
		fltSlice: slice,
	}
}
