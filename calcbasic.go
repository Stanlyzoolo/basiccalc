package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Expression struct {
	X, Y      int
	Calculate Operate
}

func main() {
	strexpr := "1+1"
	arrayofints := []int{}

	strEdited := TrimSpaces(strexpr) // Output: 1+1
	strofdigits := strings.Trim(strEdited, "+")

	strSplited := StrSplitting(strofdigits) // Output: [1 + 1]
	fmt.Println(strSplited)

	for _, value := range strSplited {
		digit, _ := strconv.Atoi(value)            // Здесь знаю, что ошибка со знаком +
		arrayofints = append(arrayofints, digit)
	}

	fmt.Println(arrayofints)

	result := Expression{
		X:         arrayofints[0],
		Y:         arrayofints[2],
		Calculate: Operate(),           // Запутался в этом месте
	}

	fmt.Println(result)
}

func TrimSpaces(strexpr string) string {
	return strings.ReplaceAll(strexpr, " ", "")
}

func StrSplitting(strEdited string) []string {
	return strings.Split(strEdited, "")
}

func (e *Expression) Addition() int {
	e.X = e.X + e.Y
	return e.X
}

func (e *Expression) Subtraction() int {
	e.X = e.X - e.Y
	return e.X
}

type Operate func(*Expression, byte) int

func getFunction() Operate {
	return func(e *Expression, operator byte) int {
		switch operator {
		case '+':
			return e.Addition()
		case '-':
			return e.Subtraction()
		}
		return 0
	}

}

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
