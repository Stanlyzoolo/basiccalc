package basiccalc

import (
	"fmt"
	"strings"

)

var singledigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

func Eval(input string) (int, error) {
	exp : mathstuff.Expression{}

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

		fn, isfn := mathstuff.Operators[s]
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
