package sentences

import (
	"Calculator/internal/sentences/errors"
)

const (
	_ = iota
	plus
	minus
	multiplication
)

// Operation Enum тип, содержащий значение математической операции
type Operation int8

func NewOperation(sign string) (Operation, error) {
	if sign == "+" {
		return plus, nil
	} else if sign == "-" {
		return minus, nil
	} else if sign == "*" {
		return multiplication, nil
	} else {
		return 0, &errors.UnsupportedSymbolError{Op: sign}
	}
}
