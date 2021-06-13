package main

import (
	"fmt"
	"strings"
)

// Structure of operands and the operator
type Expression struct {
	X, Y      int
	Operation Operate
}

// Filling Expression structure
func (exp *Expression) FillingExpression(stringarr []string) *Expression {

	// This part is for sequence "1+1"
	for _, elem := range stringarr {
		_, ok := singledigits[elem]

		if ok {
			exp.X = singledigits[stringarr[0]]
			exp.Y = singledigits[stringarr[2]]
		} else {
			exp.Operation = operators[stringarr[1]]
		}

	}
	exp.X = exp.Operation(exp.X, exp.Y)

	//This part is for sequence "2-1+  2"

	stringarrCut := stringarr[3:]

	for _, elem := range stringarrCut {
		_, ok := singledigits[elem]

		if ok {
			exp.Y = singledigits[stringarrCut[1]]
		} else {
			exp.Operation = operators[stringarrCut[0]]
		}
	}

	return exp
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
func PreparingInputSequence(condition string) []string {
	stringArr := []string{}
	conditionArr := strings.Split(condition, "")

	for _, str := range conditionArr {
		if str != " " {
			stringArr = append(stringArr, str)
		}
	}
	return stringArr
}

func main() {

	// First input
	firstInput := "1+1"
	preparedSequence := PreparingInputSequence(firstInput)
	fmt.Println("Prepared first sequence: ", PreparingInputSequence(firstInput)) // [1 + 1]

	firstExpression := Expression{}

	completeFirstExpression := firstExpression.FillingExpression(preparedSequence)

	resultOfFirstInput := &Expression{
		X:         completeFirstExpression.X,
		Y:         completeFirstExpression.Y,
		Operation: completeFirstExpression.Operation,
	}

	fmt.Println("Result of first input: ", resultOfFirstInput.X) // Output 2

	//Second input
	secondInput := "2+1  -2"
	prSequence := PreparingInputSequence(secondInput)
	fmt.Println("Prepared secind sequence: ", prSequence) //[2 + 1 - 2]

	secondExpression := Expression{}

	completeSecondExpression := secondExpression.FillingExpression(prSequence)

	resultOfSecondInput := &Expression{
		X:         completeSecondExpression.X,
		Y:         completeSecondExpression.Y,
		Operation: completeSecondExpression.Operation,
	}

	fmt.Println("Result of second input: ", resultOfSecondInput.Operation(resultOfSecondInput.X, resultOfSecondInput.Y)) // Output 1
}
