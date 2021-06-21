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
	x, y     int
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
		exp.x = arg
		exp.state = FirstArgument
		return nil
	}

	if exp.state == FirstArgWithOperator {
		exp.y = arg
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
		exp.x = exp.Evaluation(exp.x, exp.y)
		exp.state = FirstArgument
		return exp.x, nil
	}

	return 0, errors.New("invalid expression")
}
