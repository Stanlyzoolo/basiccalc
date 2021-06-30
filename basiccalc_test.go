package basiccalc_test

import (
	"testing"
	. "github.com/Stanlyzoolo/basiccalc"
)



func TestEval(t *testing.T) {
	want := 4
	var input string = "2+1 +1"

	got, err := Eval(input)
	if got != want {
		t.Error("Something went wrong", err)
	}

	input = "2+*"
	_, err2 := Eval(input)

	if err2 == nil {
		t.Error("Something went wrong", err2)
	}
}
