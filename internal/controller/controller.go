package controller

import (
	"Calculator/internal/sentences"
	"context"
	"strconv"
)

func CalculateSentence(ctx context.Context, inputSentences ...map[string]interface{}) (map[string]int64, error) {
	s := sentences.NewSentences()
	for _, sentence := range inputSentences {
		if sentence["type"].(string) == "calc" {
			var left string
			if v, ok := sentence["left"].(float64); ok {
				left = strconv.FormatFloat(v, 'f', -1, 64)
			} else if v, ok := sentence["left"].(string); ok {
				left = v
			}

			var right string
			if v, ok := sentence["right"].(float64); ok {
				right = strconv.FormatFloat(v, 'f', -1, 64)
			} else if v, ok := sentence["right"].(string); ok {
				right = v
			}
			err := s.AddCalc(
				sentence["op"].(string),
				sentence["var"].(string),
				left,
				right)
			if err != nil {
				return nil, err
			}
		} else if sentence["type"].(string) == "print" {
			s.AddPrint(sentence["var"].(string))
		}
	}

	errChan := make(chan error, 1)

	s.CalcWithContext(ctx, errChan)
	select {
	case err := <-errChan:
		if err != nil {
			return nil, err
		}
	}

	return s.Print(), nil

}
