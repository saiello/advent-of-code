package main

import (
	"testing"
)

func TestAllOperatorCombinationsWhen1(t *testing.T) {

	if len(AllOperatorCombinations(1, Add, Multiply)) != 2 {
		t.Errorf("Cartesian combinations len is wrong")
	}

	if len(AllOperatorCombinations(1, Add, Multiply, Concatenation)) != 3 {
		t.Errorf("Operator combinations len is wrong")
	}
}

func TestEvaluate(t *testing.T) {
	eq := Equation{left: 5, right: []int{2, 4}}

	if r := eq.Evaluate([]Operator{Add}); r != 6 {
		t.Errorf("Add(2,4) should be 6. But was %d", r)
	}

	if r := eq.Evaluate([]Operator{Multiply}); r != 8 {
		t.Errorf("Add(2,4) should be 8. But was %d", r)
	}
}

func TestFromImput(t *testing.T) {
	eq := Equation{left: 904, right: []int{4, 2, 727, 80, 91}}
	if !eq.isValid() {
		t.Errorf("Equation should be valid")
	}
}

func TestConcatenationAgainstExampleInput(t *testing.T) {
	eq := Equation{left: 156, right: []int{15, 6}}
	if !eq.isValid() {
		t.Errorf("Equation should be valid")
	}
}
