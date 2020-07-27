package customfile

import "github.com/stretchr/testify/mock"

type MockCustomfile struct {
	mock.Mock
}

func (m *MockCustomfile) GiveConvertedFile(fileName string) (err error) {
	args := m.Called(fileName)
	return args.Error(0)
}
