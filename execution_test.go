package expr

import (
	"errors"
	"reflect"
	"testing"
)

var tests = []struct {
	op        Operator
	x, y      any
	want      any
	wantError error
}{
	{RelationalOperatorEqualTo, 1, 1, true, nil},
	{RelationalOperatorEqualTo, 0, 0, true, nil},
	{RelationalOperatorEqualTo, -1, -1, true, nil},
	{RelationalOperatorEqualTo, 1, 2, false, nil},

	{RelationalOperatorGreaterThan, 2, 1, true, nil},
	{RelationalOperatorGreaterThan, 2, -1, true, nil},
	{RelationalOperatorGreaterThan, 1, 2, false, nil},
	{RelationalOperatorGreaterThan, -1, 2, false, nil},

	{RelationalOperatorGreaterThanOrEqualTo, 2, 1, true, nil},
	{RelationalOperatorGreaterThanOrEqualTo, 2, -1, true, nil},
	{RelationalOperatorGreaterThanOrEqualTo, 2, 2, true, nil},
	{RelationalOperatorGreaterThanOrEqualTo, 1, 2, false, nil},
	{RelationalOperatorGreaterThanOrEqualTo, -1, 2, false, nil},

	{Operator("MOCK_UNSUPPORTED_OPERATOR"), 1, 1, false, errors.New("Unsupported operator MOCK_UNSUPPORTED_OPERATOR, Can't execute the given operator MOCK_UNSUPPORTED_OPERATOR on expression (1 MOCK_UNSUPPORTED_OPERATOR 1)")},
}

func TestExecuteRelationalOperator(t *testing.T) {
	for _, test := range tests {
		var (
			got    any
			gotErr error
		)

		switch test.x.(type) {
		case int:
			got, gotErr = Execute(test.x.(int), test.op, test.y.(int))
		case string:
			got, gotErr = Execute(test.x.(string), test.op, test.y.(string))
		case float64:
			got, gotErr = Execute(test.x.(float64), test.op, test.y.(float64))
		case uintptr:
			got, gotErr = Execute(test.x.(uintptr), test.op, test.y.(uintptr))
		}

		if test.wantError == nil && !reflect.DeepEqual(got, test.want) {
			t.Errorf("Execute() = got %v, want %v", got, test.want)
		}

		if test.wantError != nil && !reflect.DeepEqual(gotErr.Error(), test.wantError.Error()) {
			t.Errorf("Execute() = got %v, want %v", gotErr, test.wantError)
		}
	}
}
