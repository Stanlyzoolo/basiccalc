package basiccalc

import (
	"fmt"
	"strings"
)

//nolint
// pickArgument provides selection of argument according to the input rune.
func pickArgument(s string) (int, bool) {
	switch s {
	case "0":
		return 0, true
	case "1":
		return 1, true
	case "2":
		return 2, true
	case "3":
		return 3, true
	case "4":
		return 4, true
	case "5":
		return 5, true
	case "6":
		return 6, true
	case "7":
		return 7, true
	case "8":
		return 8, true
	case "9":
		return 9, true
	}
	return 0, false
}

func Eval(input string) (int, error) {
	var exp expression

	result := 0

	var operatorError error

	for i, s := range strings.Split(input, "") {
		if s == " " {
			continue
		}

		arg, isDigit := pickArgument(s)

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

		fn, isfn := pickOperator(s)
		if isfn {
			operatorError := exp.SetOperator(fn)
			if operatorError != nil {
				return 0, fmt.Errorf("%w at position %v", operatorError, i)
			}

			continue
		}
	}

	return result, operatorError
}
