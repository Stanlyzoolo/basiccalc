package basiccalc

import (
	"errors"
) 

type Action func(int, int) int

var operators = map[string]Action{			
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
}

type expression struct {			
	x, y     int
	evaluation Action				
	state     int
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

	return errors.New("unexpected argument")
}

func (exp *expression) SetOperator(fn Action) error {
	if exp.state == FirstArgument {
		exp.evaluation = fn
		exp.state = FirstArgWithOperator
		return nil
	}

	return errors.New("unexpected operator")
}

func (exp *expression) Calculate() (int, error) {
	if exp.state == Ready {
		exp.x = exp.evaluation(exp.x, exp.y)
		exp.state = FirstArgument
		return exp.x, nil
	}

	return 0, errors.New("invalid expression")
}
