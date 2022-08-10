package main

import "errors"

type Calculator struct {
	num1 int
	num2 int
}

func (c *Calculator) Calculator_Add() (int, error) {
	if c.num1 < 0 || c.num2 < 0 {
		return -1, errors.New("numbers should never be negative")
	}
	return c.num1 + c.num2, nil
}
