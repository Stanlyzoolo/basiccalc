package basiccalc

import (
	"errors"
)

// Action type is a simple function for performing simple mathematical
//  operations depending on the operator
type Action func(int, int) int

// operators is a map where keys represent mathematical operators as a string
//  type and values represent the corresponding function
var operators = map[string]Action{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
}

// expression is a trivial implementation of a sequence consisting of two arguments,
//  an operator and a state of fullness of the structure
type expression struct {
	x, y       int
	evaluation Action
	state      int
}

// Constants describe states of fullness of the expression structure using
// an approach of FSM (finite state machine)
const (
	Initialized int = iota
	FirstArgument
	FirstArgWithOperator
	Ready
)

// isReady is a method to check the fullness of expression structure
// by comparing its state
func (exp *expression) IsReady() bool {
	return exp.state == Ready
}

// SetArgument is a method that  takes an argument type int, checks current
// state of the structure, assigns its value to the corresponding field
// and returns an error
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

// SetOperator is a method that takes an argument type Action, checks current
// state of the structure, assigns its value – function from operators map
// to the evaluation field and returns an error
func (exp *expression) SetOperator(fn Action) error {
	if exp.state == FirstArgument {
		exp.evaluation = fn
		exp.state = FirstArgWithOperator
		return nil
	}

	return errors.New("unexpected operator")
}

// Calculate is a method that checks the current state of the structure,
//  performs a mathematical operation according to function in the evaluation field,
// assigns its value to the first field, and returns an error
func (exp *expression) Calculate() (int, error) {
	if exp.state == Ready {
		exp.x = exp.evaluation(exp.x, exp.y)
		exp.state = FirstArgument
		return exp.x, nil
	}

	return 0, errors.New("invalid expression")
}
