package basiccalc

import (
	"errors"
)

type action func(int, int) int

// pickArgument provides selection of argument according to the input rune.
func pickOperator(s string) (action, bool) {
	if s == "+" {
		return func(x, y int) int { return x + y }, true
	}

	if s == "-" {
		return func(x, y int) int { return x - y }, true
	}

	return nil, false
}

type expression struct {
	x, y       int
	evaluation action
	state      int
}

const (
	Initialized int = iota
	FirstArgument
	FirstArgWithOperator
	Ready
)

func (exp *expression) IsReady() bool {
	return exp.state == Ready
}

var errArg = errors.New("unexpected argument")

func (exp *expression) SetArgument(arg int) error {
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

	return errArg
}

var errOp = errors.New("unexpected operator")

func (exp *expression) SetOperator(fn action) error {
	if exp.state == FirstArgument {
		exp.evaluation = fn
		exp.state = FirstArgWithOperator

		return nil
	}

	return errOp
}

var errCalc = errors.New("invalid expression")

func (exp *expression) Calculate() (int, error) {
	if exp.state == Ready {
		exp.x = exp.evaluation(exp.x, exp.y)
		exp.state = FirstArgument

		return exp.x, nil
	}

	return 0, errCalc
}
