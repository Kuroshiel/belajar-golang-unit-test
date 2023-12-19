package helper

import (
	"fmt"
	"testing"
)

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
