package basiccalc

import (
	"fmt"

)

func ExampleEval() {
	fmt.Println(Eval("1+1"))
	// Output: 2 <nil>

	// fmt.Println(Eval("2-1 + 2"))
	// Output: 3 <nil>
}