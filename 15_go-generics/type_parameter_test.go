package gogenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var result = Length("Eko")
	assert.Equal(t, "Eko", result)

	var resultNumber = Length(100)
	assert.Equal(t, 100, resultNumber)
}
