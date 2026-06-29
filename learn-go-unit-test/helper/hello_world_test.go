package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Dinar")
	}
}

func TestTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Eko)",
			request:  "Jun",
			expected: "Hello Eko",
		},
		{
			name:     "HelloWorld(Junaidi)",
			request:  "Junaidi",
			expected: "Hello Junaidi",
		},
		{
			name:     "HelloWorld(Dono)",
			request:  "Dono",
			expected: "Hello Dono",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Eko", func(t *testing.T) {
		result := HelloWorld("Eko")
		assert.Equal(t, "Hello Eko", result, "result must be: 'Hello Eko'")
		fmt.Println("TestHelloWorld done")
	})
	t.Run("Junaidi", func(t *testing.T) {
		result := HelloWorld("kong")
		assert.Equal(t, "Hello Ayu", result, "result must be: 'Hello Ayu'")
		fmt.Println("TestHelloWorld done")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum Unit Test")
	m.Run()
	fmt.Println("Setelah unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("kong")
	assert.Equal(t, "Hello Ayu", result, "result must be: 'Hello Ayu'")
	fmt.Println("TestHelloWorld done")
}

func TestHelloKanedy(t *testing.T) {
	result := HelloWorld("Kanedy")
	assert.Equal(t, "Hello Kanedy", result, "result must be: 'Hello Kanedy'")
	fmt.Println("TestHelloKhanedy done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Dono")
	require.Equal(t, "Hello Dono", result, "Result must be: 'Hello Dono'")
	fmt.Println("TestHelloWorldRequire is done")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Unit test tida bisa berjalan di linux")
	}
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result)
}
