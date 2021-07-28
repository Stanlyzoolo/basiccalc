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
	x          int
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

// token represents a type for setting arguments and operators.
type token struct {
	r rune
}

// Rune gets the value of the r field.
func (t token) rune() rune {
	return t.r
}

// tokener represents Rune() getter.
type tokener interface {
	rune() rune
}

// A tokenOperand implements operand by embedding token type.
type tokenOperand struct {
	token
	val int
}

// Value gets the value of the val field.
func (t tokenOperand) value() int {
	return t.val
}

// valuer represents Value() getter.
type valuer interface {
	value() int
}

// A tokenOperator implements operator by embedding token type.
type tokenOperator struct {
	token
	op action
}

// Operator gets the value of the op field.
func (t tokenOperator) operator() action {
	return t.op
}

// operatorer represents Operator() getter.
type operatorer interface {
	operator() action
}

// A tokenSpace implements space by embedding token type.
type tokenSpace struct {
	token
}

// isSpace reports whether the rune is a space.
func (t tokenSpace) isSpace() bool {
	return true
}

// spacer represents boolean condition.
type spacer interface {
	isSpace() bool
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

// tokenFactory returns interface depending on the incoming rune.
func tokenFactory(r rune) (tokener, error) {
	if val, ok := pickArgument(r); ok {
		return tokenOperand{token: token{r: r}, val: val}, nil
	}

	if op, ok := pickOperator(r); ok {
		return tokenOperator{token: token{r: r}, op: op}, nil
	}

	if r == ' ' {
		return tokenSpace{token: token{r: r}}, nil
	}

	return token{}, errTokenFactory
}

var errToken = errors.New("unexpected token in setToken()")

// setToken processes tokener interfaces by cheking its type
// for setting arguments and operator.
func (e *expression) setToken(t tokener) (int, error) {
	if tv, ok := t.(valuer); ok {
		return e.setArgument(tv.value())
	}

	if to, ok := t.(operatorer); ok {
		return e.setOperator(to.operator())
	}

	if _, ok := t.(spacer); ok {
		return e.x, nil
	}

	return e.x, errToken
}
