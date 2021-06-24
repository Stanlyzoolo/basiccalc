// Package basiccalc provides a primitive implementation of a basic calculator for 
// calculating simple expressions consisting of single digits and mathematical
//  addition and subtraction operators.
package basiccalc

import (
	"fmt"
	"strings"

)
// singledigits is a map where keys represent single digits
//  as a string type and values represent them in type int
var singledigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

// Eval provides evaluation of input string representing an expression 
// and returns result of mathematical operations
func Eval(input string) (int, error) {
	exp := expression{}

	result := 0
 
	var operatorError error

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

			if exp.IsReady() {
				result, _ = exp.Calculate()
			}
			continue
		}

		fn, isfn := operators[s]
		if isfn {
			operatorError := exp.SetOperator(fn)
			if operatorError != nil {
				return 0, fmt.Errorf("%s at position %v", operatorError, i)
			}
			continue
		}

	}

	return result, operatorError
}
