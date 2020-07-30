package customstrslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	methodNameToUpper          = "ToUpper"
	methodNameToLower          = "ToLower"
	methodNameLenSlice         = "LenSlice"
	methodNameGetFirstElem     = "GetFirstElem"
	methodNameReplaceFirstElem = "ReplaceFirstElem"
)

var (
	c     = []string{"one", "two", "three"}
	c1    = []string{"ONE", "TWO", "THREE"}
	expec = &customStrSlice{
		strSlice: []string{"one", "two", "three"},
	}
	expec1 = &customStrSlice{
		strSlice: []string{"ONE", "TWO", "THREE"},
	}
)

func TestCustomStrSlice_LenSlice(t *testing.T) {
	c = []string{"one", "two", "three"}
	t.Run(methodNameLenSlice, func(t *testing.T) {
		in := NewCustomStrSlice(c)
		assert.Equal(t, 3, in.LenSlice())
	})
}

func TestCustomStrSlice_GetFirstElem(t *testing.T) {
	c = []string{"one", "two", "three"}
	t.Run(methodNameGetFirstElem, func(t *testing.T) {
		in := NewCustomStrSlice(c)
		assert.Equal(t, "one", in.GetFirstElem())
	})
}

func TestCustomStrSlice_ReplaceFirstElem(t *testing.T) {
	c = []string{"one", "two", "three"}
	t.Run(methodNameReplaceFirstElem, func(t *testing.T) {
		in := NewCustomStrSlice(c)
		in.ReplaceFirstElem("lalalala")
		assert.Equal(t, "lalalala", in.GetFirstElem())
	})
}

func TestCustomStrSlice_ToLower(t *testing.T) {
	c1 = []string{"ONE", "TWO", "THREE"}
	t.Run(methodNameToLower, func(t *testing.T) {
		in := NewCustomStrSlice(c1)
		in.ToLower()
		assert.Equal(t, expec, in)
	})
}

func TestCustomStrSlice_ToUpper(t *testing.T) {
	c = []string{"one", "two", "three"}
	t.Run(methodNameToUpper, func(t *testing.T) {
		in := NewCustomStrSlice(c)
		in.ToUpper()
		assert.Equal(t, expec1, in)
	})
}
