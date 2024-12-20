package calc

import "errors"

var (
	ErrInvalidExpression = errors.New("Invalid Expression")
	ErrDivisionByZero    = errors.New("Division by Zero")
	ErrEmptyExpression   = errors.New("Empty Expression")
)
