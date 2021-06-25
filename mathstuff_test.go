package basiccalc

import (
	"errors"
	"testing"
)

func TestSetArgument(t *testing.T) {
	testTable := map[int]error{Initialized: nil, FirstArgument: errors.New("fail state FirstArgument"), FirstArgWithOperator: nil, Ready: errors.New("fail state Ready")}

	for s, e := range testTable {
		expr := expression{state: s}

		err := expr.SetArgument(2)

		if e == nil && err != nil {
			t.Error("failed SetArgument; want err = nil, got err != nil")
		}

	}

}

func TestSetOperator(t *testing.T) {

	testTable := map[int]error{Initialized: errors.New("fail state Initialized"), FirstArgument: nil, FirstArgWithOperator: errors.New("fail state FirstArgWithOperator"), Ready: errors.New("fail state Ready")}

	for s, e := range testTable {
		expr := expression{state: s}

		err := expr.SetOperator(func(int, int) int { return 0 })

		if e == nil && err != nil {
			t.Error("failed SetOperator; want err = nil, got err != nil")
		}
	}

}

func TestCalculate(t *testing.T) {

	testTable := map[int]error{Initialized: errors.New("fail state Initialized"), FirstArgument: errors.New("fail state Ready"), FirstArgWithOperator: errors.New("fail state FirstArgWithOperator"), Ready: nil}

	for s, e := range testTable {
		expr := expression{state: s}
		want := 4
		got, err := expr.Calculate()

		if e == nil && err != nil && got != want {
			t.Error("failed Calculate; want err = nil, got err != nil")
		}
	}
}
