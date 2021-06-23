package mathstuff

import (
	"testing"
)

var TestExpression = Expression{
	x:          2,
	y:          1,
	evaluation: Operators["+"],
	state:      3,
}

func TestSetArgument(t *testing.T) {

	arg := 3
	if TestExpression.state != Ready {
		err := TestExpression.SetArgument(arg)
		if err != nil {
			t.Error("State of TestExpression in not Initialized. Fail SetArgument() method")
		}
	}

	if TestExpression.state == Ready {
		err := TestExpression.SetArgument(arg)
		if err != nil {
			t.Error("State of TestExpression is Ready and already has argument")
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

func TestEvaluate(t *testing.T) {
	want := 3

	if TestExpression.state == Ready {
		got := TestExpression.evaluation(TestExpression.x, TestExpression.y)

		if want != got {
			t.Error("Wrong evaluation")
		}
	}
}
