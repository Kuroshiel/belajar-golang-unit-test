package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "Result must be 'Hello Eko'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWolrdEko(t *testing.T) {
	result := HelloWorld("Eko")

	if result != "Hello Eko" {
		//error
		//t.Fail()
		t.Error("Result must be 'Hello Eko")
	}
	fmt.Println("TestHelloWordEko Done")
}

func TestHelloWolrdKhannedy(t *testing.T) {
	result := HelloWorld("Khannedy")

	if result != "Hello Khannedy" {
		//error
		//t.FailNow()
		t.Fatal("Result must be 'Hello Khannedy")
	}
	fmt.Println("TestHelloWordKhannedy Done")
}
