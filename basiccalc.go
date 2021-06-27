// Package basiccalc provides a primitive implementation of a basic calculator for
// calculating simple expressions consisting of single digits and mathematical
//  addition and subtraction operators.
package basiccalc

import (
	"fmt"
)


// Eval provides evaluation of input string representing an expression
// and returns result of mathematical operations
func Eval(input string) (int, error) {
	exp := expression{}

	var result int
	var err error

	//  проблема большого тела цикла

	for _, r := range input {

		tk := Factory(r)

		result, err = exp.SetToken(tk)

		if err != nil {
			fmt.Println("Something went wrong")
		}
	}
	return result, err
}
