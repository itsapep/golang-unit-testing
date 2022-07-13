package main

type Calculator struct {
	num1 int
	num2 int
}

func (c *Calculator) Calculator_Add() int {
	return c.num1 + c.num2
}
