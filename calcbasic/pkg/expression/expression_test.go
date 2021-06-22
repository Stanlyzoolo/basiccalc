package expression

import (
	"fmt"

	"github.com/stretchr/testify/suite"
	// "testing"
)

type ExpressionTestSuite struct {
	suite.Suite
	expression Expression
}

func (suite *ExpressionTestSuite) TestingSetArgument(arg int) {
	arg = 2
	suite.Equal(suite.expression.x, 2)
	arg = 3
	suite.Equal(suite.expression.y, 3)
}

func (suite *ExpressionTestSuite) TestingSetOperator(fn Action) {
	fn = Operators["+"]
	suite.Equal(suite.expression.Evaluation, Operators["+"])
}

func (suite *ExpressionTestSuite) TestEvaluate() {
	expected := 5
	got := suite.expression.Evaluation(suite.expression.x, suite.expression.y)

	if expected != got {
		fmt.Println("Wrong result of evaluation")
	}
}

// var TestingStruct = Expression{
// 	x:          2,
// 	y:          1,
// 	Evaluation: Operators["+"],
// 	state:      3,
// }

// func (exp *Expression) TestSetArgument(arg int, t *testing.T) {
// 	arg = 2
// 	exp.SetArgument(arg)
// 	if exp.x != TestingStruct.x {
// 		t.Error("Arguments of x field are not equal")
// 	}
// }

// func (exp *Expression) TestSetOperator(fn Action) {
// 	operator := Operators["+"]
// 	exp.SetOperator(operator)
// 	// exp.Evaluation == fn
// 	// не понял как это тестировать
// }

// func (exp *Expression) TestEvaluate(t *testing.T) {
// 	expected := TestingStruct.Evaluation(TestingStruct.x, TestingStruct.y)
// 	got := exp.Evaluation(exp.x, 1)

// 	if expected != got {
// 		t.Error("Wrong result of evaluation")
// 	}
// }
