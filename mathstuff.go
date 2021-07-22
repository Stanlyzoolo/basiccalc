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
	x      int
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

// token represents a type for setting arguments and operators.
type token struct {
	r       rune
	variety int
}

// kind gets the value of the variety field.
func (t token) kind() int {
	return t.variety
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

// singledigits is a map where keys represent single digits
//  as a string type and values represent them in type int.
var singledigits = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

// tokenFactory returns interface depending on the incoming rune.
func tokenFactory(r rune) (tokener, error) {

	if val, ok := singledigits[r]; ok {
		return tokenOperand{token: token{r: r}, val: val}, nil
	}

	if op, ok := operators[r]; ok {
		return tokenOperator{token: token{r: r}, op: op}, nil
	}

	if r == ' ' {
		return tokenSpace{token: token{r: r}}, nil
	}

	return token{}, errors.New("unexpected token in tokenFactory()")
}

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

	return e.x, errors.New("unexpected token in setToken()")
}
