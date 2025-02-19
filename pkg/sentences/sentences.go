package sentences

import (
	"Calculator/pkg/sentences/interfaces"
	"Calculator/pkg/sentences/types"
	"strconv"
	"sync"
)

type Sentences struct {
	expressions map[*types.Variable]*types.Sentence
	variables   map[string]*types.Variable
	mu          *sync.Mutex
}

func NewSentences() *Sentences {
	return &Sentences{
		expressions: make(map[*types.Variable]*types.Sentence),
		variables:   make(map[string]*types.Variable),
		mu:          &sync.Mutex{},
	}
}

func createSentenceArgument(s string) (arg interfaces.SentenceArgument, isVar bool) {
	value, err := strconv.Atoi(s)
	if err != nil {
		return types.NewVariable(s), true
	} else {
		return types.Number(value), false
	}
}

func (s *Sentences) addVariableIfNotContainsAndSubscribe(varName string) (*interfaces.SentenceArgument, chan int64) {
	arg, isVar := createSentenceArgument(varName)
	var subscribe chan int64
	if isVar {
		if _, ok := s.variables[varName]; !ok {
			s.variables[arg.(*types.Variable).Name()] = arg.(*types.Variable)
		}
		subscribe, _ = s.variables[arg.(*types.Variable).Name()].SubscribeOnValueWhenItCalculated()
	} else {
		subscribe = nil
	}
	return &arg, subscribe
}

func (s *Sentences) AddCalc(op string, varName string, left string, right string) error {
	correctOp, err := types.NewOperation(op)
	if err != nil {
		return err
	}

	correctVar := types.NewVariable(varName)
	s.variables[correctVar.Name()] = correctVar

	leftArg, leftChan := s.addVariableIfNotContainsAndSubscribe(left)
	rightArg, rightChan := s.addVariableIfNotContainsAndSubscribe(right)

	sentence := types.NewSentence(correctVar, correctOp, leftArg, leftChan, rightArg, rightChan)

	s.mu.Lock()
	defer s.mu.Unlock()
	s.expressions[correctVar] = sentence

	return nil
}

func (s *Sentences) AddPrint(varName string) {
	if definedVar, ok := s.variables[varName]; ok {
		definedVar.MakePrintable()
	} else {
		s.variables[varName] = types.NewVariable(varName)
	}
}

func (s *Sentences) Calc() {
	wg := sync.WaitGroup{}
	for _, sentence := range s.expressions {
		errChan := make(chan error, 1)
		wg.Add(1)
		go func() {
			sentence.Calculate(errChan)
			wg.Done()
		}()
	}
	wg.Wait()
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
