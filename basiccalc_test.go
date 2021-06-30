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

	inputS := "2+*"
	_, erro := Eval(inputS)
	if erro != nil {
		t.Error("Something went wrong", erro)
	}
}
