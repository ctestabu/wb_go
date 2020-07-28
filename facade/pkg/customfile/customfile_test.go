package customfile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	methodNameGiveConvertedFile = "GiveConvertedFile"
	validInput                  = "test1.old"
)

func TestCustomFileSuccessGiveConvertedFile(t *testing.T) {
	t.Run(methodNameGiveConvertedFile, func(t *testing.T) {
		f := NewCustomFile(validInput)
		err := f.GiveConvertedFile(validInput)
		assert.NoError(t, err)
	})
}

func TestCustomFileEmptyGiveConvertedFile(t *testing.T) {
	t.Run(methodNameGiveConvertedFile, func(t *testing.T) {
		f := NewCustomFile("")
		err := f.GiveConvertedFile("")
		assert.Error(t, err)
	})
}
