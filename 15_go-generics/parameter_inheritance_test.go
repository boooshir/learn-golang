package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type VicePresident interface {
	GetName() string
	GetVicePresidentName() string
}

type MyVicePresidentName struct {
	Name string
}

func (m *MyVicePresidentName) GetName() string {
	return m.Name
}

func (m *MyVicePresidentName) GetVicePresidentName() string {
	return m.Name
}

func TestGetName(t *testing.T) {
	assert.Equal(t, "Eko", GetName[Manager](&MyManager{Name: "Eko"}))
	assert.Equal(t, "Boo", GetName[VicePresident](&MyVicePresidentName{Name: "Boo"}))
}
