package main

import (
	"fmt"
	"strconv"
)

func main () {
	Somefunc("1+1")
}

func Somefunc (s string) int {
	var result int
	strSlice := []string{}

	for _, str := range s {
		strSlice = append(strSlice, string(str))
	}

	for _, item := range strSlice {
		digit1, _ := strconv.Atoi(strSlice[0])
		digit2, _ := strconv.Atoi(strSlice[2])
		switch item {
		case "+":
			result = digit1 + digit2
		case "-":
			result = digit1 - digit2
		}
	}

	fmt.Println(result)
	return result
}