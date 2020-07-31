package visitor

import (
	"errors"
	"strings"

	v1 "github.com/wb_go/visitor/pkg/api/v1"
	"github.com/wb_go/visitor/pkg/customfloatslice"
	"github.com/wb_go/visitor/pkg/customintslice"
	"github.com/wb_go/visitor/pkg/customstrslice"
)

// Visitor ...
type Visitor interface {
	InsertInMiddle(c customintslice.CustomIntSlice) (err error)
	ReplaceStrSlice(c customstrslice.CustomStrSlice) (err error)
	RoundAndReplace(c customfloatslice.CustomFloatSlice) (err error)
}

type arr struct {
	elem int
}

// InsertInMiddle inserts value in the middle or custom slice
func (a *arr) InsertInMiddle(c customintslice.CustomIntSlice) (err error) {
	if c.LenSlice() == 0 {
		err = errors.New(v1.ErrorNotAllowed)
		return
	}
	mid := int(c.LenSlice() / 2)
	err = c.Insert(mid, a.elem)
	return
}

// ReplaceStrSlice replaces first elem with a string Replaced
func (a *arr) ReplaceStrSlice(c customstrslice.CustomStrSlice) (err error) {
	if c.LenSlice() == 0 {
		err = errors.New(v1.ErrorNotAllowed)
		return
	}
	tmp := strings.Repeat(v1.WordToReplace, a.elem)
	err = c.ReplaceFirstElem(tmp)
	return
}

// RoundAndReplace rounds and replace first elem in slice
func (a *arr) RoundAndReplace(c customfloatslice.CustomFloatSlice) (err error) {
	if c.LenSlice() == 0 {
		err = errors.New(v1.ErrorNotAllowed)
		return
	}
	tmp := float64(a.elem)
	c.ReplaceFirstElem(tmp)
	err = c.Round()
	return
}

// NewArr creates new entity of a Visitor
func NewArr(elem int) Visitor {
	return &arr{
		elem: elem,
	}
}
