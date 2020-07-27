package changer

import "github.com/stretchr/testify/mock"

type MockChanger struct {
	mock.Mock
}

func (m *MockChanger) ChangeFileName(fileName, format string) (msg string, err error) {
	args := m.Called(fileName, format)
	return args.Get(0).(string), args.Error(1)
}

func (m *MockChanger) GetFileToChange() (str string) {
	args := m.Called()
	return args.Get(0).(string)
}
