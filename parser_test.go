package main

import (
	"strings"
	"testing"
)

func TestValidParsing(t *testing.T) {
	input := strings.NewReader("TODO: Write some tests about parser\nDONE: Write parser implementation")
	parser := NewParser(input)
	expected := []Todo{
		Todo{Finished: false, Task: "Write some tests about parser"},
		Todo{Finished: true, Task: "Write parser implementation"},
	}
	result, err := parser.Parse()
	if err != nil {
		t.Errorf("Expected no error but got: %v", err)
	}
	for index, _ := range expected {
		if result[index] != expected[index] {
			t.Errorf("Expected %#v to equal %#v", result[index], expected[index])

		}
	}
}
