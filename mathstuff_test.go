package basiccalc

import (
	"errors"
	"testing"
)

func TestSetArgument(t *testing.T) {
	
	testTable := map[int]bool{
		Initialized: true, 
		FirstArgument: false,
		FirstArgWithOperator: true,
	}

	// any int value
	var arg int = 2

	for s, ok := range testTable {
		expr := &expression{state: s}

		result, err := expr.setArgument(arg)

		if ok && err != nil {
			t.Error(result, "failed SetArgument; want err = nil, got err != nil")
		}
	}
}

func TestSetOperator(t *testing.T) {

	testTable := map[int]error{Initialized: errors.New("fail state Initialized"), FirstArgument: nil, FirstArgWithOperator: errors.New("fail state FirstArgWithOperator")}

	for s, e := range testTable {
		expr := expression{state: s}

		result, err := expr.setOperator(func(int, int) int { return 0 })

		if e == nil && err != nil {
			t.Error(result, "failed SetOperator; want err = nil, got err != nil")
		}
	}

}

func TestTokenFactory(t *testing.T) {
	wantArg := tokenOperand{token: token{r: '5'}, val: 5}

	var rArg rune = '5'

	gotArg, errArg := tokenFactory(rArg)

	if gotArg != wantArg && errArg != nil {
		t.Error("unexpected operand token")
	}


	// wantOp := tokenOperator{token: token{r: '+'}, op: func(x, y int) int { return x + y }}

	// var rOp rune = '+'

	// gotOp, errOp := tokenFactory(rOp)
	// if gotOp != wantOp && errOp != nil {		//Trouble
	// 	t.Error("unexpected operator token")
	// }

	var rSpace rune = ' '
	wantSpace := tokenSpace{token: token{r: rSpace}}

	gotSpace, errSpace := tokenFactory(rSpace)

	if gotSpace != wantSpace && errSpace != nil {
		t.Error("unexpected space token")
	} 
	
}

func TestSetToken(t *testing.T) {

}
