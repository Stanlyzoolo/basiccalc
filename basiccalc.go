// Package basiccalc provides a primitive implementation of a basic calculator for
// calculating simple expressions consisting of single digits and mathematical
//  addition and subtraction operators.
package basiccalc

import (
	"fmt"

	"go.uber.org/zap"
)

// evalError wrap cause error for more context.
func evalError(cause error, p int) error {
	return fmt.Errorf("%s at position %v", cause, p)
}

// Eval provides evaluation of input string representing an expression
// and returns result of mathematical operations.
func Eval(input string) (int, error) {
	var exp expression

	var result int

	logger, _ := zap.NewDevelopment()

	for p, r := range input {

		tk, err := tokenFactory(r)
		if err != nil {
			logger.Error("failed to eval input expression",
				zap.String("package", "basiccalc"),
				zap.String("function", "tokenFactory"),
				zap.String("input", input),
				zap.Int("position", p),
				zap.Error(err),
			)
			return 0, evalError(err, p)
		}

		result, _ = exp.setToken(tk)

	}
	return result, nil
}
