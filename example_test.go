package calc_test

import (
	"fmt"
	"github.com/Stanlyzoolo/basiccalc"
)

func ExampleEval() {
	fmt.Println(calc.Eval("1+1"))
	// Output: 2

	fmt.Println(calc.Eval("2-1 + 2"))
	// Output: 3
}