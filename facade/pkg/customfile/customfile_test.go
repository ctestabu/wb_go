package customfile

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	methodNameGiveConvertedFile = "GiveConvertedFile"
	validInput = "test1.old"
	empty = ""
	invalidInput = "test1.invalid"
)

func TestCustomFileSuccessGiveConvertedFile(t *testing.T) {
	t.Run(methodNameGiveConvertedFile, func(t *testing.T) {
		f:= New(validInput)
		err := f.GiveConvertedFile(validInput)
		assert.NoError(t, err)
	})
}

func TestCustomFileEmptyGiveConvertedFile(t *testing.T) {
	t.Run(methodNameGiveConvertedFile, func(t *testing.T) {
		f:= New(empty)
		err := f.GiveConvertedFile(empty)
		assert.Error(t, err)
	})
}
