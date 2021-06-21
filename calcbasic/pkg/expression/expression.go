package expression

import "errors"

const (
	State0 int = iota
	State1
	State2
	State3
)

type Expression struct {
	X, Y      int
	Operation Operate
	state     int // 0 - New struct, 1 - X, 2 - X and Operation, 3 - X, Y and Operation
}

var singledigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

type Operate func(int, int) int

var operators = map[string]Operate{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
}

func (exp *Expression) isReady() bool {
	return exp.state == State3
}

func (exp *Expression) SetArgument(arg int) error {
	if exp.state == 0 {
		exp.X = arg
		exp.state = 1
		return nil
	}

	if exp.state == 2 {
		exp.Y = arg
		exp.state = 3
		return nil
	}

	return errors.New("unexpected argument")
}

func (exp *Expression) SetOperator(fn Operate) error {
	if exp.state == 1 {
		exp.Operation = fn
		exp.state = 2
		return nil
	}

	return errors.New("unexpected operator")
}

func (exp *Expression) Evaluate() (int, error) {
	if exp.state == 3 {
		exp.X = exp.Operation(exp.X, exp.Y)
		exp.state = 1
		return exp.X, nil
	}

	return 0, errors.New("invalid expression")
}
