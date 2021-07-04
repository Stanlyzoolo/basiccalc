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

func detectType(t token) int {
	switch t.kind {
	case Operand:
		return 0
	case Operator:
		return 1
	case Space:
		return 2
	default:
		return 3
	}
}

func TestTokenFactory(t *testing.T) {

	testTable := map[rune]token{
		'2': {r: '2', val: 2, kind: Operator},
		'+': {r: '+', op: func(x, y int) int { return x + y }, kind: Operator},
		' ': {r: ' ', kind: Space},
		'*': {},
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
	tBad := token{r: '*', kind: 4}

	_, err := expr.setToken(tBad)

	if err == nil {
		t.Error("failed tokenFactory(); want err = nil, got err != nil")
	}

}

func TestRune(t *testing.T) {
	var want rune = '1'

	tk := token{r: want}

	if tk.Rune() != want {
		t.Error("failed token.Rune()")
	}
}

func TestValue(t *testing.T) {
	var want int = 1

	tk := token{val: want}

	if tk.Value() != want {
		t.Error("failed tokenOperand.Value()")
	}
}

func TestType(t *testing.T) {
	var want int = 2

	tk := token{kind: want}

	if tk.Type() != want {
		t.Error("failed tokenOperand.Type()")
	}
}
