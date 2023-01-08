package main

import (
	"testing"
)

// move files main.go and main_test.go to src\package
// then from main directory (C:\Users\USER\go\GoTest#1)
// run with console 'go test ./...'
func TestUnitTestFramework(t *testing.T) {
	if 1 != 0 {
		t.Error("didnt work")
		t.Fail()
	}
}

func TestCache(t *testing.T) {
	insert("key1", "value1111")
	insert("key2", "value2222")
	insert("key3", "value3333")

	if len(data) != 3 {
		t.Error("expecting 3")
		t.Fail()
	}
}
