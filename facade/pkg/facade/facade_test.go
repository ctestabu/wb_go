package facade

import (
	"testing"

	"github.com/stretchr/testify/assert"

	changer2 "github.com/wb_go/facade/pkg/changer"
	"github.com/wb_go/facade/pkg/customfile"
)

const (
	methodNameReceive           = "Receive"
	methodNameChangeFileName    = "ChangeFileName"
	methodNameGetFileToChange   = "GetFileToChange"
	methodNameGiveConvertedFile = "GiveConvertedFile"
	expectedOut                 = "File to convert test1.old"
	validInput                  = "test1.old"
	validOut                    = "test1.lal"
	validFormat                 = ".lal"
)

func TestConverterSuccessReceive(t *testing.T) {
	t.Run(methodNameReceive, func(t *testing.T) {
		toChange := new(changer2.MockChanger)
		toChange.On(methodNameChangeFileName, validInput, validFormat).Return(expectedOut, nil).Once()
		toChange.On(methodNameGetFileToChange).Return(validOut).Once()

		customFile := new(customfile.MockCustomfile)
		customFile.On(methodNameGiveConvertedFile, validOut).Return(nil).Once()

		res := NewConverter(customFile, toChange)
		_, err := res.Receive(validInput, validFormat)
		assert.NoError(t, err, expectedOut, err)

	})
}
