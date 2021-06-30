package basiccalc

import (
	"errors"
)

// action type is a simple function for performing simple mathematical
//  operations depending on the operator
type action func(int, int) int

// operators is a map where keys represent mathematical operators as a string
//  type and values represent the corresponding function
var operators = map[rune]action{
	'+': func(x, y int) int { return x + y },
	'-': func(x, y int) int { return x - y },
}

// expression is a trivial implementation of a sequence consisting of two arguments,
//  an operator and a state of fullness of the structure
type expression struct {
	x, y       int
	evaluation action
	state      int
}

// Constants describe states of fullness of the expression structure using
// an approach of FSM (finite state machine)
const (
	Initialized int = iota
	FirstArgument
	FirstArgWithOperator
)

// setArgument is a method that  takes an argument type int, checks current
// state of the structure, assigns its value to the corresponding field
// and returns an error
func (exp *expression) setArgument(arg int) (int, error) {

	if exp.state == Initialized {
		exp.x = arg
		exp.state = FirstArgument
		return exp.x, nil
	}

	if exp.state == FirstArgWithOperator {
		exp.x = exp.evaluation(exp.x, arg)
		exp.state = FirstArgument
		return exp.x, nil
	}

	return exp.x, errors.New("unexpected argument")
}

// setOperator is a method that takes an argument type Action, checks current
// state of the structure, assigns its value – function from operators map
// to the evaluation field and returns an error
func (exp *expression) setOperator(fn action) (int, error) {
	if exp.state == FirstArgument {
		exp.evaluation = fn
		exp.state = FirstArgWithOperator
		return exp.x, nil
	}

	return exp.x, errors.New("unexpected operator")
}

type token struct {
	r    rune
	kind int
}

// Getter...
func (t token) Type() int {
	return t.kind
}

// Getter...
func (t token) Rune() rune {
	return t.r
}

type tokener interface {
	Rune() rune
}

type tokenOperand struct {
	token
	val int
}

// Getter...
func (t tokenOperand) Value() int {
	return t.val
}

type valuer interface {
	Value() int
}

type tokenOperator struct {
	token
	op action
}

// Getter...
func (t tokenOperator) Operator() action {
	return t.op
}

type operatorer interface {
	Operator() action
}

type tokenSpace struct {
	token
}

// Compare method
func (t tokenSpace) isSpace() bool {
	return true
}

type spacer interface {
	isSpace() bool
}

// singledigits is a map where keys represent single digits
//  as a string type and values represent them in type int
var singledigits = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

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

	return token{}, errors.New("unexpected token")
}

func (exp *expression) setToken(t tokener) (int, error) {

	if tv, ok := t.(valuer); ok {
		return exp.setArgument(tv.Value())
	}

	if to, ok := t.(operatorer); ok {
		return exp.setOperator(to.Operator())
	}

	if _, ok := t.(spacer); ok {
		return exp.x, nil
	}

	return exp.x, errors.New("unexpected token")
}
