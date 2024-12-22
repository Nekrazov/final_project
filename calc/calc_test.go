package calc_test

import (
	"testing"

	"github.com/IvanSolomatin/calc_go/pkg/calculation"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "Сложение",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "Приоритет скобок",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "Деление",
			expression:     "1/4",
			expectedResult: 0,25,
		},
		{
			name:           "Приоритет операций",
			expression:     "3+3*3",
			expectedResult: 12,
		},
        {
            name:           "Комбинация",
            expression:     "1+2+3*5/6",
            expectedResult: 5,
        }
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calc.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression) 
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult) 
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:        "invalid_operator",
			expression:  "1+1*",
			expectedErr: calc.ErrInvalidExpression,
		},
		{
			name:        "invalid_operator",
			expression:  "2+2**2",
			expectedErr: calc.ErrInvalidExpression,
		},
		{
			name:        "invalid_parentheses",
			expression:  "((2+2-*(2",
			expectedErr: calc.ErrInvalidExpression,
		},
		{
			name:        "empty_expression",
			expression:  "",
			expectedErr: calc.ErrEmptyExpression,
		},
		{
			name:        "division_by_zero",
			expression:  "1/0",
			expectedErr: calc.ErrDivisionByZero,
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calc.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is invalid but result %f was obtained", testCase.expression, val) 
			}

			if err != testCase.expectedErr {
				t.Fatalf("expected error %v, got %v", testCase.expectedErr, err) 
			}
		})
	}
}