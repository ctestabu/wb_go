package customintslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	methodNameReverse  = "ReversSlice"
	methodNameInsert   = "Insert"
	methodNameDelete   = "Delete"
	methodNameLenSlice = "LenSlice"
)

var (
	a      = []int{1, 2, 7, 3, 22}
	expec  = &customIntSlice{intSlice: []int{2, 7, 3, 22}}
	expec1 = &customIntSlice{intSlice: []int{777, 1, 2, 7, 3, 22}}
	expec2 = &customIntSlice{intSlice: []int{22, 3, 7, 2, 1}}
)

func TestCustomIntSlice_LenSlice(t *testing.T) {
	a = []int{1, 2, 7, 3, 22}
	t.Run(methodNameLenSlice, func(t *testing.T) {
		in := NewCustomIntSlice(a)
		assert.Equal(t, 5, in.LenSlice())
	})
}

func TestCustomIntSlice_Delete(t *testing.T) {
	a = []int{1, 2, 7, 3, 22}
	t.Run(methodNameDelete, func(t *testing.T) {
		in := NewCustomIntSlice(a)
		in.Delete(0)
		tmp := in
		assert.Equal(t, expec, tmp)
	})
}

func TestCustomIntSlice_Insert(t *testing.T) {
	a = []int{1, 2, 7, 3, 22}
	t.Run(methodNameInsert, func(t *testing.T) {
		in := NewCustomIntSlice(a)
		in.Insert(0, 777)
		tmp := in
		assert.Equal(t, expec1, tmp)

	})
}

func TestCustomIntSlice_ReversSlice(t *testing.T) {
	a = []int{1, 2, 7, 3, 22}
	t.Run(methodNameReverse, func(t *testing.T) {
		in := NewCustomIntSlice(a)
		in.ReversSlice()
		tmp := in
		assert.Equal(t, expec2, tmp)
	})
}
