package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("via")

	if result != "Hello via" {
		//error
		t.Error("Result must be 'Hello Via'") //t.T t.Fail or recomended : t.Fatal  t.Error
	}
	fmt.Println("test done")
}
func TestHelloWorldVeaa(t *testing.T) {
	result := HelloWorld("vea")

	if result != "Hello vea" {
		//error
		t.Fatal("Result must be 'Hello Vea'")
	}
}
