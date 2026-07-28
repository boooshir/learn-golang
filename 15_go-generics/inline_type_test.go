package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int | int16 | int64 | float64 }](first T, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, int64(100), FindMin(int64(100), int64(200)))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))
}

func GetFirst[T []E, E any](data T) E {
	first := data[0]

	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Eko", "Kurniawan", "Kanedi",
	}

	first := GetFirst(names)
	assert.Equal(t, "Eko", first)
}
