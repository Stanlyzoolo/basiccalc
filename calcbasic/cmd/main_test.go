package main

import (
	"testing"
)

var input string = "2+1"

func TestCalculate(t *testing.T) {
	expected := 3

	result, err := Calculate(input)
	if result != expected {
		t.Error("Something went wrong", err)
	}
}
