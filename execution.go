package expr

import (
	"cmp"
	"fmt"
)

type Operator string

const (
	RelationalOperatorEqualTo              Operator = "EQUAL_TO"
	RelationalOperatorGreaterThan          Operator = "GREATER_THAN"
	RelationalOperatorGreaterThanOrEqualTo Operator = "GREATER_THAN_OR_EQUAL_TO"

	// INFO: Guideline for another expression operator
	// ArithmeticOperatorPlus                 Operator = "PLUS"
	// ArithmeticOperatorMinus                Operator = "MINUS"
	// ArithmeticOperatorTimes                Operator = "TIMES"
	// ArithmeticOperatorDivide               Operator = "DIVIDE"
	// ...
	// RelationalOperatorLessThan             Operator = "LESS_THAN"
	// RelationalOperatorLessThanOrEqualTo    Operator = "LESS_THAN_OR_EQUAL_TO"
	// RelationalOperatorIncludes             Operator = "INCLUDES"
	// ...
)

func Execute[OP1 cmp.Ordered, OP2 cmp.Ordered](operand1 OP1, operator Operator, operand2 OP1) (any, error) {
	executers := map[Operator]func() any{
		RelationalOperatorEqualTo:              func() any { return cmp.Compare(operand1, operand2) == 0 },
		RelationalOperatorGreaterThan:          func() any { return cmp.Compare(operand1, operand2) == +1 },
		RelationalOperatorGreaterThanOrEqualTo: func() any { return !cmp.Less(operand1, operand2) },
	}

	executer, found := executers[operator]
	if !found {
		return nil, fmt.Errorf("Unsupported operator %v, Can't execute the given operator %v on expression (%v %v %v)", operator, operator, operand1, operator, operand2)
	}

	return executer(), nil
}
