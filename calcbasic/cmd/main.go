package main

import (
	"fmt"
	"strings"
	
)

type EvalError struct {
	Message  string
	Position int
	err      error
}

func (e *EvalError) Error() string {
	return fmt.Sprintf("%s at position: %d", e.Message, e.Position)
}

func Calculate(input string) (int, error) {
	exp := Expression{}

	result := 0
	var newError error

	for i, s := range strings.Split(input, "") {

		if s == " " {
			continue
		}

		arg, isDigit := singledigits[s]

		if isDigit {

			err := exp.SetArgument(arg)
			if err != nil {
				return 0, err
			}

			if exp.isReady() {
				result, err = exp.Evaluate()
				if err != nil {
					return 0, err
				}
				return result, nil
			}

			continue
		}

		fn, isfn := operators[s]
		if isfn {
			newError := exp.SetOperator(fn)
			if newError != nil {
				return 0, &EvalError{"invalid expression", i, newError}
			}
			continue
		}

	}

	return result, newError
}
