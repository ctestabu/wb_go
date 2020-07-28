package customfile

import "github.com/stretchr/testify/mock"

// MockCustomfile ...
type MockCustomfile struct {
	mock.Mock
}

// GiveConvertedFile ...
func (m *MockCustomfile) GiveConvertedFile(fileName string) (err error) {
	args := m.Called(fileName)
	return args.Error(0)
}
