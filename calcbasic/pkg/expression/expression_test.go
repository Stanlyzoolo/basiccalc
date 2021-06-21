package expression

import (
	"testing"
)

var TestingStruct = Expression{
	x:          2,
	y:          1,
	Evaluation: Operators["+"],
	state:      3,
}

func (exp *Expression) TestSetArgument(arg int, t *testing.T) {
	arg = 2
	exp.SetArgument(arg)
	if exp.x != TestingStruct.x {
		t.Error("Arguments of x field are not equal")
	}
}

func (exp *Expression) TestSetOperator(fn Action) {
	operator := Operators["+"]
	exp.SetOperator(operator)
	// не понял как это тестировать
}

func (exp *Expression) TestEvaluate(t *testing.T) {
	expected := TestingStruct.Evaluation(TestingStruct.x, TestingStruct.y)
	got := exp.Evaluation(exp.x, 1)

	if expected != got {
		t.Error("Wrong result of evaluation")
	}
}
