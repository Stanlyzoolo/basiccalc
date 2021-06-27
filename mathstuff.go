package basiccalc

import (
	"errors"
	"unicode"
)

// Action type is a simple function for performing simple mathematical
//  operations depending on the operator
type Action func(int, int) int

// operators is a map where keys represent mathematical operators as a string
//  type and values represent the corresponding function
var operators = map[rune]Action{
	'+': func(x, y int) int { return x + y },
	'-': func(x, y int) int { return x - y },
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

func (exp *expression) SetToken(tk token) (int, error) {

	// Начинается проверка типа токена

	err := exp.SetArgument(FirstArgWithOperator)
	if err != nil {
		return 0, err
	}

	// if r == ' ' {
	// 	continue
	// }
	// arg, isDigit := singledigits[r]

	// if isDigit {

	// 	err := exp.SetArgument(arg)
	// 	if err != nil {
	// 		return 0, err
	// 	}

	// 	if exp.IsReady() {
	// 		result, _ = exp.Calculate()
	// 	}
	// 	continue
	// }

	// fn, isfn := operators[r]

	// if isfn {
	// 	operatorError := exp.SetOperator(fn)
	// 	if operatorError != nil {
	// 		return 0, fmt.Errorf("%s at position %v", operatorError, i)
	// 	}
	// 	continue
	// }
}

// singledigits is a map where keys represent single digits
//  as a string type and values represent them in type int
var singledigits = map[rune]int{
	'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
}

type token struct {
	r   rune
	val int
	op  Action
}

func (t token) Type() {

	// определиться с поведением и что возвращает

	// здесь пока не разобрался
}


// Интерфейс для token
type Tokener interface {
	Operand()
	Operation()
	Empty()
}

// Здесь реализовал методы, которые четко определяют, в моем понимании, какой токен
// возвращается из TokenFactory

// Далее, соответствующие выходные токены надо передать в SetArgument и SetOperator

func (t token) Operand(val int) token {
	return token{val: val}
}

func (t token) Operation(op Action) token {
	return token{op: op}
}

func (t token) Empty(r rune) token {
	return token{r: r}
}

func Factory(r rune) token {
	var t token

	val, ok := singledigits[r]

	if ok {
		return t.Operand(val)
	}

	op, ok := operators[r]

	if ok {
		return t.Operation(op)
	}

	// просто вернуть пустой токен или тип - пустой токен
	if unicode.IsSpace(r) {
		return t.Empty(r)
	}
	return t

}

// почитать про токены

// почитать статью https://ruslanspivak.com/lsbasi-part1/
