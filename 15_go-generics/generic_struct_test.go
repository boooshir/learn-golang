package gogenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangFirst(first T) T {
	d.First = first
	return first
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "boshir",
		Second: "Ahmad",
	}
	assert.Equal(t, "budi", data.ChangFirst("budi"))
	assert.Equal(t, "Hello su", data.SayHello("su"))
	fmt.Println(data)
}
