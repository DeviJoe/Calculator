package main

import (
	"Calculator/pkg/sentences"
	"fmt"
)

func main() {
	s := sentences.NewSentences()
	err := s.AddCalc("+", "x", "10", "2")
	if err != nil {
		panic(err)
	}
	err = s.AddCalc("*", "y", "x", "5")
	if err != nil {
		panic(err)
	}

	s.AddPrint("x")
	s.AddPrint("y")

	s.AddPrint("x")
	s.Calc()

	results := s.Print()
	for varName, value := range results {
		fmt.Printf("VAR: %s, VALUE IS %d\n", varName, value)
	}
}
