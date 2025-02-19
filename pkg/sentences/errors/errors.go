package errors

import "fmt"

type UnsupportedSymbolError struct {
	Op string
}

func (e *UnsupportedSymbolError) Error() string {
	return fmt.Sprintf(
		"unsupported symbol: not supported operand %s",
		e.Op,
	)
}

type ResettingVariableValueError struct {
	VariableName       string
	OldValue           int64
	TryingSettingValue int64
}

func (e *ResettingVariableValueError) Error() string {
	return fmt.Sprintf(
		"Trying resetting %d value for variable %s with it value %d (varaible alrady setted)",
		e.TryingSettingValue,
		e.VariableName,
		e.OldValue,
	)
}

type CannotSubscribeToVariableError struct {
	VariableName int64
}

func (e *CannotSubscribeToVariableError) Error() string {
	return fmt.Sprintf("Cannot subscribe to variable %d", e.VariableName)
}
