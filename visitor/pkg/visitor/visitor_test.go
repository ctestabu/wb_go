package visitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wb_go/visitor/pkg/customfloatslice"
	"github.com/wb_go/visitor/pkg/customintslice"
	"github.com/wb_go/visitor/pkg/customstrslice"
)

const (
	methodNameInsertInMiddle  = "InsertInMiddle"
	methodNameReplaceStrSlice = "ReplaceStrSlice"
	methodNameRoundAndSwap    = "RoundAndSwap"
)

var (
	a        = []int{1, 2, 7, 3, 22}
	aBad     []int
	b        = []float64{1.1, 2.2, 3.4343, 5.5}
	bBad     []float64
	c        = []string{"one", "two", "three"}
	cBad     []string
	test     = 4
	expecInt = customintslice.NewCustomIntSlice([]int{1, 2, 4, 7, 3, 22})
	expecStr = customstrslice.NewCustomStrSlice([]string{"ReplacedReplacedReplacedReplaced", "two", "three"})
	expecFlt = customfloatslice.NewCustomFloatSlice([]float64{4, 2, 3, 6})
)

func TestArr_InsertInMiddle(t *testing.T) {
	a = []int{1, 2, 7, 3, 22}
	t.Run(methodNameInsertInMiddle, func(t *testing.T) {
		in := customintslice.NewCustomIntSlice(a)
		val := NewArr(test)
		err := val.InsertInMiddle(in)
		assert.Equal(t, expecInt, in)
		assert.NoError(t, err)
	})
}
func TestArr_InsertInMiddleBad(t *testing.T) {
	t.Run(methodNameInsertInMiddle, func(t *testing.T) {
		in := customintslice.NewCustomIntSlice(aBad)
		val := NewArr(test)
		re2 := val.InsertInMiddle(in)
		assert.Error(t, re2)
	})
}

func TestArr_ReplaceStrSlice(t *testing.T) {
	c = []string{"one", "two", "three"}
	t.Run(methodNameReplaceStrSlice, func(t *testing.T) {
		in := customstrslice.NewCustomStrSlice(c)
		val := NewArr(test)
		err := val.ReplaceStrSlice(in)
		assert.Equal(t, expecStr, in)
		assert.NoError(t, err)
	})
}

func TestArr_ReplaceStrSliceBad(t *testing.T) {
	t.Run(methodNameReplaceStrSlice, func(t *testing.T) {
		in := customstrslice.NewCustomStrSlice(cBad)
		val := NewArr(test)
		err := val.ReplaceStrSlice(in)
		assert.Error(t, err)
	})
}

func TestArr_RoundAndReplace(t *testing.T) {
	b = []float64{1.1, 2.2, 3.4343, 5.5}
	t.Run(methodNameRoundAndSwap, func(t *testing.T) {
		in := customfloatslice.NewCustomFloatSlice(b)
		val := NewArr(test)
		err := val.RoundAndReplace(in)
		assert.Equal(t, expecFlt, in)
		assert.NoError(t, err)
	})
}

func TestArr_RoundAndReplaceBad(t *testing.T) {
	t.Run(methodNameRoundAndSwap, func(t *testing.T) {
		in := customfloatslice.NewCustomFloatSlice(bBad)
		val := NewArr(test)
		err := val.RoundAndReplace(in)
		assert.Error(t, err)
	})
}
