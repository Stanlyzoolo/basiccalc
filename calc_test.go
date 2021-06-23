package calc_test

import (
	"testing"
)

var input string = "2+1"

func TestEval(t *testing.T) {
	want := 3

	got, err := Eval(input)
	if got != want {
		t.Error("Something went wrong", err)
	}
}
