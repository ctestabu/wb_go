package changer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	methodNameChangeFileName = "ChangeFileName"
	validInput               = "test1.old"
	invalidInput             = "test1.invalid"
)

func TestChangerCorrectChangeFileName(t *testing.T) {
	t.Run(methodNameChangeFileName, func(t *testing.T) {
		f := NewChanger(validInput)
		_, err := f.ChangeFileName(validInput, ".lal")
		assert.NoError(t, err)

	})
}

func TestChangerIncorrectChangeFileName(t *testing.T) {
	t.Run(methodNameChangeFileName, func(t *testing.T) {
		f := NewChanger(invalidInput)
		_, err := f.ChangeFileName(invalidInput, ".invalid")
		assert.Error(t, err)
	})

}
