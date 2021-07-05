package basiccalc

import (
	"errors"
	"testing"
)

func TestSetArgument(t *testing.T) {

	testTable := map[int]bool{
		Initialized:          true,
		FirstArgument:        false,
		FirstArgWithOperator: true,
	}

	// any int value
	var arg int = 2

	expr := expression{
		evaluation: func(x, y int) int { return x + y },
	}

	for s, ok := range testTable {
		expr.state = s
		result, err := expr.setArgument(arg)

		if ok && err != nil {
			t.Error(result, "failed SetArgument(); want err = nil, got err != nil")
		}
	}
}

func TestSetOperator(t *testing.T) {

	testTable := map[int]error{
		Initialized:          errors.New("fail state Initialized"),
		FirstArgument:        nil,
		FirstArgWithOperator: errors.New("fail state FirstArgWithOperator"),
	}

	for s, e := range testTable {
		expr := expression{state: s}

		result, err := expr.setOperator(func(int, int) int { return 0 })

		if e == nil && err != nil {
			t.Error(result, "failed SetOperator(); want err = nil, got err != nil")
		}
	}

}

func detectType(t tokener) int {
	switch t.(type) {
	case tokenOperand:
		return 0
	case tokenOperator:
		return 1
	case tokenSpace:
		return 2
	default:
		return 3
	}
}

func TestTokenFactory(t *testing.T) {

	testTable := map[rune]tokener{
		'2': tokenOperand{token: token{r: '2'}, val: 2},
		'+': tokenOperator{token: token{r: '+'}, op: func(x, y int) int { return x + y }},
		' ': tokenSpace{token: token{r: ' '}},
		'*': token{},
	}

	for r, want := range testTable {

		got, err := tokenFactory(r)

		if detectType(got) != detectType(want) && err != nil {
			t.Error("failed tokenFactory(); want err = nil, got err != nil")
		}
	}
}

func TestSetToken(t *testing.T) {
	expr := expression{}
	tBad := token{r: '*'}

	_, err := expr.setToken(tBad)

	if err == nil {
		t.Error("failed tokenFactory(); want err = nil, got err != nil")
	}

}

func TestRune(t *testing.T) {
	var want rune = '1'

	tk := token{r: want}

	if tk.rune() != want {
		t.Error("failed token.Rune()")
	}
}

func TestValue(t *testing.T) {
	var want int = 1

	tk := tokenOperand{val: want}

	if tk.value() != want {
		t.Error("failed tokenOperand.Value()")
	}
}

func TestType(t *testing.T) {
	var want int = 2

	tk := token{variety: want}

	if tk.kind() != want {
		t.Error("failed tokenOperand.Type()")
	}
}
