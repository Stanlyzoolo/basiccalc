package basiccalc

import (
	"errors"
)

// action represents a simple function for performing simple mathematical
//  operations depending on the operator.
type action func(int, int) int

// pickArgument provides selection of argument according to the input rune.
func pickOperator(r rune) (action, bool) {
	if r == '+' {
		return func(x, y int) int { return x + y }, true
	}

	if r == '-' {
		return func(x, y int) int { return x - y }, true
	}

	return nil, false
}

// expression represents a trivial implementation of a sequence consisting of two arguments,
// an operator and a state of fullness of the structure.
type expression struct {
	x, y          int
	evaluation action
	state      int
}

// Constants describe states of fullness of the expression structure using
// an approach of FSM (finite state machine).
const (
	Initialized int = iota
	FirstArgument
	FirstArgWithOperator
)

var errArg = errors.New("unexpected argument")

// setArgument takes an argument, checks current
// state of the structure and assigns its value to the corresponding field.
func (e *expression) setArgument(arg int) (int, error) {
	if e.state == Initialized {
		e.x = arg
		e.state = FirstArgument

		return e.x, nil
	}

	if e.state == FirstArgWithOperator {
		e.x = e.evaluation(e.x, arg)
		e.state = FirstArgument

		return e.x, nil
	}

	return e.x, errArg
}

var errOp = errors.New("unexpected operator")

// setOperator takes action type function, checks current
// state of the structure and assigns its value – function from operators map
// to the evaluation field.
func (e *expression) setOperator(fn action) (int, error) {
	if e.state == FirstArgument {
		e.evaluation = fn
		e.state = FirstArgWithOperator

		return e.x, nil
	}

	return e.x, errOp
}

// constants describe variety of arguments for the expression.
const (
	Operand int = iota
	Operator
	Space
)

// token represents a type for setting arguments and operators.
type token struct {
	r       rune
	val     int
	op      action
	variety int
}

// value gets the value of the val field.
func (t token) value() int {
	return t.val
}

// operator gets the value of the op field.
func (t token) operator() action {
	return t.op
}

// isSpace reports whether the rune is a space.
func isSpace(r rune) bool {
	return r == ' '
}

//nolint
// pickArgument provides selection of argument according to the input rune.
func pickArgument(r rune) (int, bool) {
	switch r {
	case '0':
		return 0, true
	case '1':
		return 1, true
	case '2':
		return 2, true
	case '3':
		return 3, true
	case '4':
		return 4, true
	case '5':
		return 5, true
	case '6':
		return 6, true
	case '7':
		return 7, true
	case '8':
		return 8, true
	case '9':
		return 9, true
	}
	return 0, false
}

var errTokenFactory = errors.New("unexpected token in tokenFactory")

// tokenFactory returns token depending on the incoming rune.
func tokenFactory(r rune) (token, error) {
	if val, ok := pickArgument(r); ok {
		return token{r: r, val: val, variety: Operand, op: nil}, nil
	}

	if op, ok := pickOperator(r); ok {
		return token{r: r, op: op, variety: Operator, val: 0}, nil
	}

	if isSpace(r) {
		return token{r: r, variety: Space, val: 0, op: nil}, nil
	}

	return token{}, errTokenFactory
}

var errToken = errors.New("unexpected token in setToken()")

// setToken processes tokens by cheking its type
// for setting arguments and operator.
func (e *expression) setToken(t token) (int, error) {
	if t.variety == Operand {
		return e.setArgument(t.value())
	}

	if t.variety == Operator {
		return e.setOperator(t.operator())
	}

	if t.variety == Space {
		return e.x, nil
	}

	return e.x, errToken
}
