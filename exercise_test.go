package main

import (
	"reflect"
	"testing"
)

var strexpr string = "5-1 -  3"

func Test_TrimSpaces(t *testing.T) {
	var result string = "5-1-3"
	if result!= TrimSpaces(strexpr){
		t.Errorf("Result is not %v", result)
	}
}


func Test_StrSplitting(t *testing.T) {
	var strEdited string = TrimSpaces(strexpr)
	var strSplited []string = StrSplitting(strEdited)
	result := StrSplitting(strEdited)
	if reflect.DeepEqual(strSplited, result) != true {
		t.Errorf("Result is not %v", result)
	}
}

func Test_Addition(t *testing.T) {
	var num1, num2 int = 1, 1
	var result int = 2
	if result != Addition(num1, num2) {
		t.Errorf("Result is not %v", result)
	}
}

func Test_Subtraction(t *testing.T) {
	var num1, num2 int = 1, 1
	var result int = 4
	if result != Subtraction(num1, num2) {
		t.Errorf("Result is not %v", result)
	}
}

func Test_Multiplication(t *testing.T) {
	var num1, num2 int = 1, 1
	var result int = 5
	if result != Multiplication(num1, num2) {
		t.Errorf("Result is not %v", result)
	}
}

func Test_Division(t *testing.T) {
	var num1, num2 int = 1, 1
	var result int = 5
	if result != Division(num1, num2) {
		t.Errorf("Result is not %v", result)
	}
}