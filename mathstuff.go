package basiccalc

import (
	"errors"
)

// action represents a simple function for performing simple mathematical
//  operations depending on the operator.
type action func(int, int) int

// operators is a map where keys represent mathematical operators as a string
//  type and values represent the corresponding function.
var operators = map[rune]action{
	'+': func(x, y int) int { return x + y },
	'-': func(x, y int) int { return x - y },
}

// expression represents a trivial implementation of a sequence consisting of two arguments,
// an operator and a state of fullness of the structure.
type expression struct {
	x, y       int
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

	return e.x, errors.New("unexpected argument")
}

// setOperator takes action type function, checks current
// state of the structure and assigns its value – function from operators map
// to the evaluation field.
func (e *expression) setOperator(fn action) (int, error) {
	if e.state == FirstArgument {
		e.evaluation = fn
		e.state = FirstArgWithOperator
		return e.x, nil
	}

	return e.x, errors.New("unexpected operator")
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

// rune gets the value of the r field.
func (t token) rune() rune {
	return t.r
}

// value gets the value of the val field.
func (t token) value() int {
	return t.val
}

// operator gets the value of the op field.
func (t token) operator() action {
	return t.op
}

// kind gets the value of the variety field.
func (t token) kind() int {
	return t.variety
}

// isSpace reports whether the rune is a space.
func isSpace(r rune) bool {
	return r == ' '
}

// singledigits is a map where keys represent single digits
//  as a string type and values represent them in type int.
var singledigits = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

// tokenFactory returns token depending on the incoming rune.
func tokenFactory(r rune) (token, error) {

	if val, ok := singledigits[r]; ok {
		return token{r: r, val: val, variety: Operand}, nil
	}

	if op, ok := operators[r]; ok {
		return token{r: r, op: op, variety: Operator}, nil
	}

	if isSpace(r) {
		return token{r: r, variety: Space}, nil
	}

	return token{}, errors.New("unexpected token in tokenFactory")
}

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

	return e.x, errors.New("unexpected token in setToken()")
}
