package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// var Test = Calculator{
// 	num1: 1,
// 	num2: 2,
// }

type CalculatorTestSuite struct {
	suite.Suite
	calcService Calculator
}

func (suite *CalculatorTestSuite) SetupTest() {
	suite.calcService = Calculator{}
}

// func (suite *CalculatorTestSuite) TestCalculator_Add(t *testing.T) {
func (suite *CalculatorTestSuite) TestCalculator_Add() {
	// 1+1=2
	// assert.Equal(t, 3, Test.Calculator_Add(), "Seharusnya 3")

	testTable := []struct {
		num1 int
		num2 int
		res  int
	}{
		{4, 4, 8},
		{-4, 5, -1},
	}
	for _, table := range testTable {
		suite.calcService.num1 = table.num1
		suite.calcService.num2 = table.num2
		actual, err := suite.calcService.Calculator_Add()
		if err != nil {
			assert.NotNil(suite.T(), err, "")
			assert.Equal(suite.T(), table.res, actual, "")
		} else {
			assert.Nil(suite.T(), err, "")
			assert.Equal(suite.T(), table.res, actual, "")
		}
		// expected := 3
		// assert.Equal(suite.T(), expected, actual, "Seharusnya 3")
	}
}

func TestCalculatorTestSuite(t *testing.T) {
	suite.Run(t, new(CalculatorTestSuite))
}
