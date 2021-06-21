package expression

import (
	"errors"
)

type Action func(int, int) int

var Operators = map[string]Action{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
}

type Expression struct {
	X, Y      int
	Evaluation Action
	state     int
}

const (
	Initialized int = iota
	FirstArgument
	FirstArgWithOperator
	Ready
)

func (exp *Expression) IsReady() bool {
	return exp.state == Ready
}

func (exp *Expression) SetArgument(arg int) error {
	if exp.state == Initialized {
		exp.X = arg
		exp.state = FirstArgument
		return nil
	}

	if exp.state == FirstArgWithOperator {
		exp.Y = arg
		exp.state = Ready
		return nil
	}

	return errors.New("unexpected argument")
}

func (exp *Expression) SetOperator(fn Action) error {
	if exp.state == FirstArgument {
		exp.Evaluation = fn
		exp.state = FirstArgWithOperator
		return nil
	}

	return errors.New("unexpected operator")
}

func (exp *Expression) Evaluate() (int, error) {
	if exp.state == Ready {
		exp.X = exp.Evaluation(exp.X, exp.Y)
		exp.state = FirstArgument
		return exp.X, nil
	}

	return 0, errors.New("invalid expression")
}
