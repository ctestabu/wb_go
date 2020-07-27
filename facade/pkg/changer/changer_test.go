package changer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	list          = []string{"test1.old", "test2.old", "test3.old"}
)
const (
	methodNameChangeFileName = "ChangeFileName"
	validInput = "test1.old"
	invalidInput = "test1.invalid"
)

func TestChangerCorrectChangeFileName(t *testing.T) {
	t.Run(methodNameChangeFileName, func(t *testing.T) {
		f:= New(validInput)
		_, err := f.ChangeFileName(validInput, ".lal")
		assert.NoError(t, err)

	})
}

func TestChangerIncorrectChangeFileName(t *testing.T) {
	t.Run(methodNameChangeFileName, func(t *testing.T) {
		f:= New(invalidInput)
		_, err := f.ChangeFileName(invalidInput, ".invalid")
		assert.Error(t, err)
	})

}
