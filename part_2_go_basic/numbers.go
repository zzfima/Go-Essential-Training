package main

import "fmt"

//CalculatorInterface describe simple calculator operation
type CalculatorInterface interface {
	AddNumbers() int
	SubNumbers() int
}

//CalculatorIntegers integers types
type CalculatorIntegers struct {
	num1, num2 int
}

//AddNumbers add two numbers
func (c CalculatorIntegers) AddNumbers() int {
	return c.num1 + c.num2
}

//SubNumbers add two numbers
func (c CalculatorIntegers) SubNumbers() int {
	return c.num1 - c.num2
}

//Measure test func
func Measure(c CalculatorInterface) {
	fmt.Println(c.AddNumbers())
	fmt.Println(c.SubNumbers())
}

func main() {
	c := CalculatorIntegers{num1: 6, num2: 2}
	Measure(c)
}
