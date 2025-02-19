package types

import (
	"Calculator/pkg/sentences/errors"
	"Calculator/pkg/sentences/interfaces"
	"strconv"
)

type Sentence struct {
	variable   *Variable
	op         Operation
	left       interfaces.SentenceArgument
	leftResCh  chan int64
	right      interfaces.SentenceArgument
	rightResCh chan int64
}

func NewSentence(
	variable *Variable,
	op Operation,
	left *interfaces.SentenceArgument,
	leftResCh chan int64,
	right *interfaces.SentenceArgument,
	rightResCh chan int64,
) *Sentence {
	return &Sentence{
		variable:   variable,
		op:         op,
		left:       *left,
		right:      *right,
		leftResCh:  leftResCh,
		rightResCh: rightResCh,
	}
}

func getValueFromVariable(arg interfaces.SentenceArgument, resCh chan int64) (value int64, e error) {
	if val, ok := arg.GetValueIsItSetting(); ok {
		return val, nil
	} else {
		return <-resCh, nil
	}
}

func (s *Sentence) Calculate(errChan chan error) {
	leftValue, err := getValueFromVariable(s.left.(interfaces.SentenceArgument), s.leftResCh)
	if err != nil {
		errChan <- err
		return
	}
	rightValue, err := getValueFromVariable(s.right.(interfaces.SentenceArgument), s.rightResCh)
	if err != nil {
		errChan <- err
		return
	}

	var value int64
	if s.op == plus {
		value = leftValue + rightValue
		err = s.variable.SetValue(value)
		if err != nil {
			errChan <- err
			return
		}
	} else if s.op == minus {
		value = leftValue - rightValue
		err = s.variable.SetValue(value)
		if err != nil {
			errChan <- err
			return
		}
	} else if s.op == multiplication {
		value = leftValue * rightValue
		err = s.variable.SetValue(value)
		if err != nil {
			errChan <- err
			return
		}
	} else {
		errChan <- &errors.UnsupportedSymbolError{Op: strconv.Itoa(int(s.op))}
		return
	}

	s.variable.mu.Lock()
	defer s.variable.mu.Unlock()
	for i := range len(s.variable.notifyChannels) {
		s.variable.notifyChannels[i] <- value
	}
	return
}

func (s *Sentence) GetVariable() *Variable {
	return s.variable
}
