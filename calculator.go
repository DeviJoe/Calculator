package main

import (
	"Calculator/internal/sentences"
	"context"
	"fmt"
	"log"
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

	err = s.AddCalc("-", "q", "y", "20")
	if err != nil {
		panic(err)
	}

	err = s.AddCalc("+", "unusedA", "y", "100")
	if err != nil {
		panic(err)
	}

	err = s.AddCalc("*", "unusedB", "unusedA", "2")
	if err != nil {
		panic(err)
	}

	s.AddPrint("q")

	err = s.AddCalc("-", "z", "x", "15")
	if err != nil {
		panic(err)
	}

	s.AddPrint("z")

	err = s.AddCalc("+", "ignoreC", "z", "y")
	if err != nil {
		panic(err)
	}

	s.AddPrint("x")

	s.AddPrint("x")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errChan := make(chan error, 1)

	s.CalcWithContext(ctx, errChan)
	select {
	case err := <-errChan:
		if err != nil {
			log.Println(err)
		}
	}

	results := s.Print()
	for varName, value := range results {
		fmt.Printf("VAR: %s, VALUE IS %d\n", varName, value)
	}
}
