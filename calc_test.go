package main

import (
	"testing"
)

var Test = Calculator{
	num1: 1,
	num2: 2,
}

func TestCalculator_Add(t *testing.T) {
	// 1+1=2
	// t.Logf("hasil:%d", Test.Calculator_Add())
	// if Test.Calculator_Add() != 3 {
	// 	t.Errorf("SALAH! harusnya %d", 3)
	// }

	t.Run("", func(t *testing.T) {
		test := Calculator{
			num1: 1,
			num2: 2,
		}
		if test.Calculator_Add() != 3 {
			t.Errorf("SALAH! harusnya %d", 3)
		}
	})
}

func TestMain(t *testing.T) {
	main()
}
