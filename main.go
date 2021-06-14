package main

import (
	"errors"
	"fmt"
	"strings"
)

// Structure of operands and the operator
type Expression struct {
	X, Y      int
	Operation Operate
	state     int // 0 - New struct, 1 - X, 2 - Operation, 3 - Y
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

func Calculate(input string) (int, error) {
	exp := Expression{}

	result := 0
	var newError  error

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

			result, newError = exp.Evaluate()
			if newError != nil {
				// Выдать ошибку с указанием индекса (место происхождения ошибки)
			}
			continue
		}

		fn, isfn := operators[s]
		if isfn {
			err := exp.SetOperator(fn)
			if err != nil {
				return 0, err
			}
			continue
		}
	}
	return result, newError
}






// Filling Expression structure
func (exp *Expression) FillingExpression(stringarr []string) (*Expression, []string) {

	for _, elem := range stringarr[:3] {
		_, ok := singledigits[elem]

		if ok {
			exp.X = singledigits[stringarr[0]]
			exp.Y = singledigits[stringarr[2]]
		} else {
			exp.Operation = operators[stringarr[1]]
		}

	}
	exp.X = exp.Operation(exp.X, exp.Y)

	stringarr = stringarr[3:]

	return exp, stringarr
}

// Operation function
type Operate func(int, int) int

// Map of operators "+" "-" and funcs
var operators = map[string]Operate{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
}

// Map of single digits
var singledigits = map[string]int{
	"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9,
}

// Preparing input condition with trim spaces
func PreparingInputSequence(condition string) []string { // Лишняя функция
	stringArr := []string{}

	for _, str := range strings.Split(condition, "") {
		if str == " " {
			continue
		}
		stringArr = append(stringArr, str)
	}
	return stringArr
}

//  Operations with filled structures
func Operations(exp *Expression, stringarr []string) int {

	for _, elem := range stringarr {
		if len(stringarr) >= 2 {

			_, ok := singledigits[elem]
			if ok {
				exp.Y = singledigits[elem]

				exp.X = exp.Operation(exp.X, exp.Y)
				stringarr = stringarr[2:]

			} else {
				exp.Operation = operators[elem]
			}

		} else {
			fmt.Println("The sequence is empty")
			break
		}

	}

	return exp.X
}

func main() {

	// First input
	firstInput := "1+1"
	preparedSequence := PreparingInputSequence(firstInput)
	fmt.Println("Prepared first sequence: ", PreparingInputSequence(firstInput)) // [1 + 1]

	firstExpression := Expression{}

	completeFirstExpression, firstSeq := firstExpression.FillingExpression(preparedSequence)

	resultOfFirstInput := Operations(completeFirstExpression, firstSeq)

	fmt.Println("Result of first input: ", resultOfFirstInput) // Output 2

	//Second input
	secondInput := "2+1  -2"
	prSequence := PreparingInputSequence(secondInput)
	fmt.Println("Prepared second sequence: ", prSequence) //[2 + 1 - 2]

	secondExpression := Expression{}

	completeSecondExpression, secondSeq := secondExpression.FillingExpression(prSequence)

	resultOfSecondInput := Operations(completeSecondExpression, secondSeq) // Output 1

	fmt.Println("Result of second input: ", resultOfSecondInput)

	// Third input
	thirdInput := "1+9-4+4-8+8-8+6"

	preSequence := PreparingInputSequence(thirdInput)
	fmt.Println("Prepared third sequence: ", preSequence)

	thirdExpression := Expression{}

	completedThirdExpression, thirdSeq := thirdExpression.FillingExpression(preSequence)

	resultOfThirdInput := Operations(completedThirdExpression, thirdSeq)

	fmt.Println("Result of third input: ", resultOfThirdInput) // Output 8

}
