package main

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/basverweij/aoc2017/day25/turing"
)

func input(s string) *machine {
	input := antlr.NewInputStream(s)
	lexer := turing.NewTuringLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := turing.NewTuringParser(stream)
	p.BuildParseTrees = true

	blueprint := p.Blueprint()
	listener := &turingListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, blueprint)

	return listener.machine
}

type turingListener struct {
	*turing.BaseTuringListener

	machine *machine
	state   *state
}

func (l *turingListener) EnterBlueprint(ctx *turing.BlueprintContext) {
	l.machine = newMachine(
		int(ctx.
			StartState().(*turing.StartStateContext).
			State().
			GetText()[0]-'A'),
		mustInt(strconv.Atoi(ctx.
			Diagnostics().(*turing.DiagnosticsContext).
			Num().
			GetText())),
	)
}

func (l *turingListener) EnterState(ctx *turing.StateContext) {
	l.state = &state{}
	l.machine.states = append(l.machine.states, l.state)
}

func (l *turingListener) EnterRules(ctx *turing.RulesContext) {
	rules := &rules{
		write: uint(mustInt(strconv.Atoi(ctx.
			Write().(*turing.WriteContext).
			Num().
			GetText()))),
		move: directionNames[ctx.
			Move().(*turing.MoveContext).
			Dir().
			GetText()],
		nextState: int(ctx.
			NextState().(*turing.NextStateContext).
			State().
			GetText()[0] - 'A'),
	}

	l.state.rules = append(l.state.rules, rules)
}

var directionNames = map[string]direction{
	"left":  left,
	"right": right,
}

func mustInt(a int, err error) int {
	if err != nil {
		panic(err)
	}

	return a
}
