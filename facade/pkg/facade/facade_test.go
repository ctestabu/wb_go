package facade

import "testing"

const (
	methodNameReceive = "Receive"
	methodNameChangeFileName = "ChangeFileName"
	methodNameGetFileToChange = "GetFileToChange"
	methodNameGiveConvertedFile = "GiveConvertedFile"


	validInput = "test1.old"
	empty = ""
	invalidInput = "test1.invalid"
)

func TestConverterSuccessReceive(t *testing.T) {
	t.Run(methodNameReceive, func(t *testing.T){

	})
}
