package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	strexpr := "5-1 -  3"
	strEdited := TrimSpaces(strexpr)      // Output: 5-1-3
	strSplited := StrSplitting(strEdited) // Output: [5 - 1 - 3]
	resultCalc := Calculate(strSplited)   // Output: 4
	fmt.Println(resultCalc)
}

func TrimSpaces(strexpr string) string {
	return strings.ReplaceAll(strexpr, " ", "")
}

func StrSplitting(strEdited string) []string {
	return strings.Split(strEdited, "")
}

func Calculate(strSplited []string) int {
	var result int
	for k := range strSplited {
		if strSplited[k+1] == "+" {
			num1, _ := strconv.Atoi(strSplited[0])
			num2, _ := strconv.Atoi(strSplited[2])
			result = Addition(num1, num2)
		}
		if strSplited[k+1] == "-" {
			num1, _ := strconv.Atoi(strSplited[0])
			num2, _ := strconv.Atoi(strSplited[2])
			result = Subtraction(num1, num2)
		} else {
			break
		}
	}
	return result
}

func Addition(num1, num2 int) int {
	return num1 + num2
}

func Subtraction(num1, num2 int) int {
	return num1 - num2
}

func Multiplication(num1, num2 int) int {
	return num1 * num2
}

func Division(num1, num2 int) int {
	return num1 / num2
}
