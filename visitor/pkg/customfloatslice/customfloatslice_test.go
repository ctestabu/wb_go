package customfloatslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	methodNameRound            = "Round"
	methodNameSwapFirstAndLast = "SwapFirstAndLast"
	methodNameGetFirstElem     = "GetFirstElem"
	methodNameReplaceFirstElem = "ReplaceFirstElem"
	methodNameLenSlice         = "LenSlice"
)

var (
	b     = []float64{1.1, 2.2, 3.4343, 5.5}
	expec = &customFloatSlice{
		fltSlice: []float64{1, 2, 3, 6},
	}
	expec1 = &customFloatSlice{
		fltSlice: []float64{5.5, 2.2, 3.4343, 1.1},
	}
)

func TestCustomFloatSlice_LenSlice(t *testing.T) {
	b = []float64{1.1, 2.2, 3.4343, 5.5}
	t.Run(methodNameLenSlice, func(t *testing.T) {
		in := NewCustomFloatSlice(b)
		assert.Equal(t, len(b), in.LenSlice())
	})
}

func TestCustomFloatSlice_Round(t *testing.T) {
	b = []float64{1.1, 2.2, 3.4343, 5.5}
	t.Run(methodNameRound, func(t *testing.T) {
		in := NewCustomFloatSlice(b)
		in.Round()
		tmp := in
		assert.Equal(t, expec, tmp)
	})
}

func TestCustomFloatSlice_SwapFirstAndLast(t *testing.T) {
	b = []float64{1.1, 2.2, 3.4343, 5.5}
	t.Run(methodNameSwapFirstAndLast, func(t *testing.T) {
		in := NewCustomFloatSlice(b)
		in.SwapFirstAndLast()
		tmp := in
		assert.Equal(t, expec1, tmp)
	})
}

func TestCustomFloatSlice_GetFirstElem(t *testing.T) {
	b = []float64{1.1, 2.2, 3.4343, 5.5}
	t.Run(methodNameGetFirstElem, func(t *testing.T) {
		in := NewCustomFloatSlice(b)
		assert.Equal(t, 1.1, in.GetFirstElem())
	})
}

func TestCustomFloatSlice_ReplaceFirstElem(t *testing.T) {
	b = []float64{1.1, 2.2, 3.4343, 5.5}
	t.Run(methodNameReplaceFirstElem, func(t *testing.T) {
		in := NewCustomFloatSlice(b)
		in.ReplaceFirstElem(8.8)
		assert.Equal(t, 8.8, in.GetFirstElem())
	})
}
