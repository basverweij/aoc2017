package main

import (
	"math"
)

type registers struct {
	values map[string]int
	maxReg int
}

type operatorFunc func(int, int) bool

var operators = map[string]operatorFunc{
	"<":  func(register, amount int) bool { return register < amount },
	"<=": func(register, amount int) bool { return register <= amount },
	">":  func(register, amount int) bool { return register > amount },
	">=": func(register, amount int) bool { return register >= amount },
	"==": func(register, amount int) bool { return register == amount },
	"!=": func(register, amount int) bool { return register != amount },
}

func newRegisters() *registers {
	return &registers{
		values: make(map[string]int),
		maxReg: math.MinInt32,
	}
}

func (r *registers) apply(i *instruction) {
	if !operators[i.ConditionOperator](r.values[i.ConditionRegister], i.ConditionAmount) {
		return
	}

	amount := i.Amount
	if !i.IsIncrease {
		amount = -amount
	}

	r.values[i.Register] += amount

	max := r.largestValue()
	if max > r.maxReg {
		r.maxReg = max
	}
}

func (r *registers) largestValue() int {
	max := math.MinInt32

	for _, v := range r.values {
		if v > max {
			max = v
		}
	}

	return max
}

func (r *registers) largestValueEver() int {
	return r.maxReg
}
