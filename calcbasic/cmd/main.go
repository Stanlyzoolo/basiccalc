package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Stanlyzoolo/homeworks/calcbasic/pkg/expression"
	"github.com/pkg/errors"
)

var singledigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

func Calculate(input string) (int, error) {
	exp := expression.Expression{}

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
				result, err = exp.Evaluate()
				if err != nil {
					return 0, err
				}
				return result, nil
			}

			continue
		}

		fn, isfn := expression.Operators[s]
		if isfn {
			operatorError := exp.SetOperator(fn)
			if operatorError != nil {
				return 0, errors.Wrapf(operatorError, "Invalid expression at position: %v", i)
			}
			continue
		}

	}

	return result, operatorError
}

func main() {
	CLIArguments := os.Args
	StringExpression := strings.Join(CLIArguments[1:],"")
	

	result, error := Calculate(StringExpression)
	fmt.Println(result, error)

}
