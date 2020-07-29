package visitor

import (
	"errors"
	"github.com/wb_go/visitor/pkg/customfloatslice"
	"github.com/wb_go/visitor/pkg/customintslice"
	"github.com/wb_go/visitor/pkg/customstrslice"
	"strings"
)

type Visitor interface {
	InsertInMiddle(c customintslice.CustomIntSlice) (err error)
	ReplaceStrSlice(c customstrslice.CustomStrSlice) (err error)
	RoundAndSwap(c customfloatslice.CustomFloatSlice) (err error)
}

type arr struct {
	elem int
}

func (a *arr) InsertInMiddle(c customintslice.CustomIntSlice) (err error) {
	if c.LenSlice() == 0 {
		err = errors.New("you are not allowed to insert in empty slice here")
		return
	}
	mid := int(c.LenSlice() / 2)
	c.Insert(mid, a.elem)
	return
}

func (a *arr) ReplaceStrSlice(c customstrslice.CustomStrSlice) (err error) {
	if c.LenSlice() == 0 {
		err = errors.New("you are not allowed to insert in empty slice here")
		return
	}
	tmp := strings.Repeat("Replaced", a.elem)
	c.ReplaceFirstElem(tmp)
	return
}

func (a *arr) RoundAndSwap(c customfloatslice.CustomFloatSlice) (err error) {
	if c.LenSlice() == 0 {
		err = errors.New("you are not allowed to insert in empty slice here")
		return
	}
	tmp := float64(a.elem)
	c.ReplaceFirstElem(tmp)
	c.Round()
	return
}

func NewArr(elem int) Visitor {
	return &arr{
		elem: elem,
	}

}
