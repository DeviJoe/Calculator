package types

import (
	"Calculator/pkg/sentences/errors"
	"sync"
)

type Variable struct {
	name           string
	value          int64
	isValueSetting bool
	isPrintable    bool
	mu             *sync.Mutex
	// Слайс каналов, в которые надо послать результат вычисления операций
	notifyChannels []chan int64
}

func NewVariable(name string) *Variable {
	return &Variable{
		name:           name,
		value:          0,
		isValueSetting: false,
		isPrintable:    false,
		mu:             &sync.Mutex{},
		notifyChannels: make([]chan int64, 0),
	}
}

func (v *Variable) Name() string {
	return v.name
}

func (v *Variable) GetValueIsItSetting() (value int64, isValueSetting bool) {
	v.mu.Lock()
	defer v.mu.Unlock()
	if v.isValueSetting {
		return v.value, v.isValueSetting
	} else {
		return 0, v.isValueSetting
	}
}

func (v *Variable) SetValue(value int64) error {
	v.mu.Lock()
	defer v.mu.Unlock()

	if !v.isValueSetting {
		v.value = value
		v.isValueSetting = true
		return nil
	} else {
		return &errors.ResettingVariableValueError{
			VariableName:       v.name,
			TryingSettingValue: value,
			OldValue:           v.value,
		}
	}
}

func (v *Variable) SubscribeOnValueWhenItCalculated() (result chan int64, isSubscribed bool) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if v.isValueSetting {
		return nil, false
	} else {
		result = make(chan int64, 1)
		v.notifyChannels = append(v.notifyChannels, result)
		return result, true
	}
}

func (v *Variable) MakePrintable() {
	v.mu.Lock()
	defer v.mu.Unlock()
	v.isPrintable = true
}

func (v *Variable) IsPrintable() bool {
	return v.isPrintable
}
