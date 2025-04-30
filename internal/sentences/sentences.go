package sentences

import (
	"Calculator/internal/sentences/errors"
	"Calculator/internal/sentences/interfaces"
	"context"
	"strconv"
	"sync"
)

type Sentences struct {
	expressions map[string]*Sentence
	variables   map[string]*Variable
	mu          *sync.Mutex
}

func NewSentences() *Sentences {
	return &Sentences{
		expressions: make(map[string]*Sentence),
		variables:   make(map[string]*Variable),
		mu:          &sync.Mutex{},
	}
}

func createSentenceArgument(s string) (arg interfaces.SentenceArgument, isVar bool) {
	value, err := strconv.Atoi(s)
	if err != nil {
		return newVariable(s), true
	} else {
		return Number(value), false
	}
}

func (s *Sentences) addVariableIfNotContainsAndSubscribe(varName string) (*interfaces.SentenceArgument, chan int64) {
	arg, isVar := createSentenceArgument(varName)
	var subscribe chan int64
	if isVar {
		if _, ok := s.variables[varName]; !ok {
			s.variables[arg.(*Variable).Name()] = arg.(*Variable)
		}
		subscribe, _ = s.variables[arg.(*Variable).Name()].SubscribeOnValueWhenItCalculated()
	} else {
		subscribe = nil
	}
	return &arg, subscribe
}

func (s *Sentences) AddCalc(op string, varName string, left string, right string) error {
	correctOp, err := NewOperation(op)
	if err != nil {
		return err
	}

	var correctVar *Variable
	if v, ok := s.variables[varName]; ok {
		if v.isValueSetting {
			return &errors.ResettingVariableValueError{VariableName: varName}
		} else {
			correctVar = v
		}
	} else {
		correctVar = newVariable(varName)
		s.variables[correctVar.Name()] = correctVar
	}
	leftArg, leftChan := s.addVariableIfNotContainsAndSubscribe(left)
	rightArg, rightChan := s.addVariableIfNotContainsAndSubscribe(right)

	sentence := newSentence(correctVar, correctOp, leftArg, leftChan, rightArg, rightChan)

	s.expressions[correctVar.Name()] = sentence

	return nil
}

func (s *Sentences) AddPrint(varName string) {
	if definedVar, ok := s.variables[varName]; ok {
		definedVar.MakePrintable()
	} else {
		s.variables[varName] = newVariable(varName)
		s.variables[varName].MakePrintable()
	}
}

func (s *Sentences) getUsedSentences() (result map[string]*Sentence) {
	result = make(map[string]*Sentence)

	for _, v := range s.variables {
		if v.isPrintable {
			result[v.name] = s.expressions[v.name]
		}
	}
	return result
}

func (s *Sentences) calc(ctx context.Context, errChan chan error) {
	wg := sync.WaitGroup{}
	for _, sentence := range s.expressions {
		wg.Add(1)
		go func() {
			resErr := make(chan error, 1)
			sentence.calculate(resErr)
			select {
			case err := <-resErr:
				if err != nil {
					errChan <- err
				}
			case <-ctx.Done():
			}
			wg.Done()
		}()
	}
	wg.Wait()
	errChan <- nil
	return
}

func (s *Sentences) CalcWithContext(ctx context.Context, errChan chan error) {
	s.calc(ctx, errChan)
}

func (s *Sentences) Print() map[string]int64 {

	result := make(map[string]int64)
	for varName, variable := range s.variables {
		if variable.IsPrintable() {
			value, isSetting := variable.GetValueIsItSetting()
			if isSetting {
				result[varName] = value
			}
		}
	}
	return result
}
