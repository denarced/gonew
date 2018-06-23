package main

import (
	"sort"
	"testing"
)

func TestGenerateDirPaths(t *testing.T) {
	// EXERCISE
	actual := generateDirPaths("pname")

	// VERIFY
	expected := []string{
		"pname/cmd/pname",
		"pname/configs",
		"pname/internal/app/pname",
		"pname/internal/pkg",
		"pname/pkg",
		"pname/scripts",
	}
	sort.Strings(actual)
	sort.Strings(expected)

	actualLength := len(actual)
	expectedLength := len(expected)
	if actualLength != expectedLength {
		t.Fatalf(
			"Dir slice length mismatch: actual %v vs expected %v",
			actualLength,
			expectedLength)
	}
	for i := 0; i < actualLength; i++ {
		eachActual := actual[i]
		eachExpected := expected[i]
		if eachActual != eachExpected {
			t.Errorf(
				"Mismatching dirs: actual %v vs expected %v",
				eachActual,
				eachExpected)
		}
	}
}
