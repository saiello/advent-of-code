package main

import "testing"

func TestNoGear(t *testing.T) {
	var sample string = `467..114..
	...*......`

	_, sol2 := Solve([]byte(sample))

	if sol2 != 0 {
		t.Errorf("Unexpected")
	}
}

func TestSameLineGear(t *testing.T) {
	var sample string = `10*3..`

	_, sol2 := Solve([]byte(sample))

	if sol2 != 30 {
		t.Errorf("Fail expected: %d but actual is %d", 30, sol2)
	}
}
