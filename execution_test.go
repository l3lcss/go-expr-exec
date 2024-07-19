package expr

import (
	"errors"
	"math"
	"reflect"
	"testing"
)

func TestExecuteRelationalOperator(t *testing.T) {
	type args struct {
		operand1 int
		operator Operator
		operand2 int
	}

	tests := []struct {
		name      string
		args      args
		want      any
		wantError error
	}{
		{
			name: "scenario: successfully executed",
			args: args{
				operand1: 2,
				operator: RelationalOperatorGreaterThanOrEqualTo,
				operand2: 1,
			},
			want: true,
		},
		{
			name: "scenario: successfully executed",
			args: args{
				operand1: 2,
				operator: RelationalOperatorGreaterThanOrEqualTo,
				operand2: 2,
			},
			want: true,
		},
		{
			name: "scenario: successfully executed",
			args: args{
				operand1: int(math.Inf(1)),
				operator: RelationalOperatorGreaterThanOrEqualTo,
				operand2: int(math.Inf(-1)),
			},
			want: true,
		},
		{
			name: "scenario: successfully executed",
			args: args{
				operand1: 2,
				operator: RelationalOperatorGreaterThanOrEqualTo,
				operand2: 3,
			},
			want: false,
		},
		{
			name: "scenario: successfully executed",
			args: args{
				operand1: int(math.NaN()),
				operator: RelationalOperatorGreaterThanOrEqualTo,
				operand2: 3,
			},
			want: false,
		},
		{
			name: "scenario: unsuccessfully executed",
			args: args{
				operand1: 2,
				operator: Operator("MOCK_UNSUPPORTED_OPERATOR"),
				operand2: 1,
			},
			want:      false,
			wantError: errors.New("Unsupported operator MOCK_UNSUPPORTED_OPERATOR, Can't execute the given operator MOCK_UNSUPPORTED_OPERATOR on expression (2 MOCK_UNSUPPORTED_OPERATOR 1)"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Execute[int, int](tt.args.operand1, tt.args.operator, tt.args.operand2)

			if tt.wantError == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Execute() = %v, want %v", got, tt.want)
			}

			if tt.wantError != nil && !reflect.DeepEqual(err.Error(), tt.wantError.Error()) {
				t.Errorf("Execute() = %v, want %v", err, tt.wantError)
			}
		})
	}
}
