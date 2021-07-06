package basiccalc

import (
	"testing"
)

var TestExpression = expression{
	x:          0,
	y:          0,
	evaluation: nil,
	state:      0,
}

func TestSetArgument(t *testing.T) {

	arg := 2
	if TestExpression.state != Ready {
		err := TestExpression.SetArgument(arg)
		if err != nil {
			t.Error("state of TestExpression in not Initialized. Fail SetArgument() method")
		}
	}

	if TestExpression.state == Ready {
		err := TestExpression.SetArgument(arg)
		if err != nil {
			t.Error("state of TestExpression is Ready and already has argument")
		}
	}

}

func TestSetOperator(t *testing.T) {
	var fn Action

	if TestExpression.state != FirstArgument {
		err := TestExpression.SetOperator(fn)
		if err != nil {
			t.Error("State of TestExpression is not FirstArgument and already has function")
		}
	}
}

func TestCalculate(t *testing.T) {
	want := 3

	if TestExpression.state == Ready {
		got := TestExpression.evaluation(TestExpression.x, TestExpression.y)

		if want != got {
			t.Error("Wrong evaluation")
		}
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
