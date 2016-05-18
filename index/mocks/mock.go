package mocks
import (
	"github.com/stretchr/testify/mock"
	"github.com/csimplestring/go-mem-store/index"
	"github.com/csimplestring/go-mem-store/document"
)

type MockedIndex struct {
	mock.Mock
}

func (m *MockedIndex) Search(op index.Op, args ...interface{}) ([]document.ObjectID, error)  {
	called := m.Called(op, args)
	return called.Get(0).([]document.ObjectID), called.Error(1)
}

func (m *MockedIndex) Field() string {
	called := m.Called()
	return called.String(0)
}
