package main

import "math"

type registers map[string]int

type operatorFunc func(int, int) bool

var operators = map[string]operatorFunc{
	"<":  func(register, amount int) bool { return register < amount },
	"<=": func(register, amount int) bool { return register <= amount },
	">":  func(register, amount int) bool { return register > amount },
	">=": func(register, amount int) bool { return register >= amount },
	"==": func(register, amount int) bool { return register == amount },
	"!=": func(register, amount int) bool { return register != amount },
}

func (r registers) apply(i *instruction) {
	if !operators[i.ConditionOperator](r[i.ConditionRegister], i.ConditionAmount) {
		return
	}

	amount := i.Amount
	if !i.IsIncrease {
		amount = -amount
	}

	r[i.Register] += amount
}

func (r registers) largestValue() int {
	max := math.MinInt32

	for _, v := range r {
		if v > max {
			max = v
		}
	}

	return max
}
