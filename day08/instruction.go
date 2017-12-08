package main

import (
	"errors"
	"regexp"
	"strconv"
)

type instruction struct {
	Register          string
	IsIncrease        bool
	Amount            int
	ConditionRegister string
	ConditionOperator string
	ConditionAmount   int
}

var (
	instructionRegexp = regexp.MustCompile("([a-z]+) (inc|dec) (-?[0-9]+) if ([a-z]+) (>|<|>=|<=|==|!=) (-?[0-9]+)")
	errInvalidSource  = errors.New("invalid source")
)

func newInstructionFromSource(s string) (*instruction, error) {
	m := instructionRegexp.FindStringSubmatch(s)
	if m == nil {
		return nil, errInvalidSource
	}

	amount, _ := strconv.Atoi(m[3])
	conditionAmount, _ := strconv.Atoi(m[6])

	return &instruction{
		Register:          m[1],
		IsIncrease:        m[2] == "inc",
		Amount:            amount,
		ConditionRegister: m[4],
		ConditionOperator: m[5],
		ConditionAmount:   conditionAmount,
	}, nil
}
