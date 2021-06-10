package main

import (
	"fmt"
	"strings"
)

type Expression struct {
	X, Y      int
	// Calculate Operate
}

func (e Expression) Additions() int {
	e.X = e.X + e.Y
	return e.X
}

func (e Expression) Subtractions() int {
	e.X = e.X - e.Y
	return e.X
}

type Operate func(int, int) int

func (op Operate) Addition(e *Expression) int {
	e.X = e.X + e.Y
	return e.X
} 

func (op Operate) Subtraction(e *Expression) int {
	e.X = e.X - e.Y
	return e.X
}

var singledigits = map[string]int {
	"0":0, "1":1, "2": 2, "3": 3, "4":4, "5": 5, "6": 6, "7": 7, "8": 8, "9":9,
}

var operators = map[string]interface{}{
	"+": Operate.Addition,
	"-": Operate.Subtraction,	
}


func main() {
	strexpr := "1+1"

	strEdited := TrimSpaces(strexpr) // Output: 1+1
	strofdigits := strings.Trim(strEdited, "+")

	strSplited := StrSplitting(strofdigits) // Output: [1 + 1]
	fmt.Println("String array: ",strSplited)

	exp := Expression{}

	for _, symbol := range strSplited {
		_, ok := singledigits[symbol]
		if ok {
			exp.X = singledigits[symbol]
		} 
			exp.Y = singledigits[symbol]
		
	}
	fmt.Println("Filled structure: ", exp)
	
}

// Working with input expresson
// Trim all spaces
func TrimSpaces(strexpr string) string {
	return strings.ReplaceAll(strexpr, " ", "")
}

// Split expression into array []string
func StrSplitting(strEdited string) []string {
	return strings.Split(strEdited, "")
}

























// func getFunction() Operate {
// 	return func(e *Expression, operator byte) int {
// 		switch operator {
// 		case '+':
// 			return e.Addition()
// 		case '-':
// 			return e.Subtraction()
// 		}
// 		return 0
// 	}

// }

// func Calculate(strSplited []string) int {
// 	var result int

// 	for k := range strSplited {
// 		operator := strSplited[k+1]
// 		if operator == "+" {
// 			num1, _ := strconv.Atoi(strSplited[0])
// 			num2, _ := strconv.Atoi(strSplited[2])
// 			result = Addition(num1, num2)
// 		}
// 		if operator == "-" {
// 			num1, _ := strconv.Atoi(strSplited[0])
// 			num2, _ := strconv.Atoi(strSplited[2])
// 			result = Subtraction(num1, num2)
// 		} else {
// 			break
// 		}
// 	}
// 	return result
// }

// func Addition(num1, num2 int) int {
// 	return num1 + num2
// }

// func Subtraction(num1, num2 int) int {
// 	return num1 - num2
// }

// func Multiplication(num1, num2 int) int {
// 	return num1 * num2
// }

// func Division(num1, num2 int) int {
// 	return num1 / num2
// }

// for _, value := range strSplited {
	// 	_, ok := operators[value]
	// 	if ok {
	// 		exp.X = operators.
	// 	}
	// }