package basiccalc_test

import (
	"testing"

	. "github.com/stanlyzoolo/basiccalc"
)

func TestEval(t *testing.T) {
	want := 4
	var input string = "2+1 +1"

	got, err := Eval(input)
	if got != want {
		t.Errorf("failed Eval() with input  '%s'; %s of input expression; want err = nil, got err != nil", input, err)
	}

	input = "2+*"
	_, err2 := Eval(input)

	if err2 == nil {
		t.Errorf("failed Eval() with input  '%s'; %s of input expression; want err = nil, got err != nil", input, err2)
	}
}

func BenchmarkEval3(b *testing.B) {
	input := "2+1"
	for n := 0; n < b.N; n++ {
		Eval(input)
	}
}

func BenchmarkEva10(b *testing.B) {
	input := "2+1   +1-8"
	for n := 0; n < b.N; n++ {
		Eval(input)
	}
}

func BenchmarkEval30(b *testing.B) {
	input := "2+1-1   +8  -4+3 -1+2 +3-8+9+5"
	for n := 0; n < b.N; n++ {
		Eval(input)
	}
}
