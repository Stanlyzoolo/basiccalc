// Package basiccalc provides a primitive implementation of a basic calculator for
// calculating simple expressions consisting of single digits and mathematical
//  addition and subtraction operators.
package basiccalc

import (
	"fmt"
)

// evalError wrap cause error for more context.
func evalError(cause error, p int) error {
	return fmt.Errorf("%s at position %v", cause, p)
}

// Eval provides evaluation of input string representing an expression
// and returns result of mathematical operations.
func Eval(input string) (int, error) {
	var exp expression = expression{}

	var result int

	for p, r := range input {

		tk, err := tokenFactory(r)
		if err != nil {
			return 0, evalError(err, p)
		}

		result, _ = exp.setToken(tk)

	}
	return result, nil
}
