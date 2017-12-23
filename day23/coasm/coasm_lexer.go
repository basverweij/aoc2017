// Code generated from .\CoAsm.g4 by ANTLR 4.7.1. DO NOT EDIT.

package coasm

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 6, 42, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 24, 10, 2, 3,
	3, 3, 3, 3, 4, 5, 4, 29, 10, 4, 3, 4, 6, 4, 32, 10, 4, 13, 4, 14, 4, 33,
	3, 5, 6, 5, 37, 10, 5, 13, 5, 14, 5, 38, 3, 5, 3, 5, 2, 2, 6, 3, 3, 5,
	4, 7, 5, 9, 6, 3, 2, 4, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2,
	47, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2,
	2, 3, 23, 3, 2, 2, 2, 5, 25, 3, 2, 2, 2, 7, 28, 3, 2, 2, 2, 9, 36, 3, 2,
	2, 2, 11, 12, 7, 117, 2, 2, 12, 13, 7, 103, 2, 2, 13, 24, 7, 118, 2, 2,
	14, 15, 7, 117, 2, 2, 15, 16, 7, 119, 2, 2, 16, 24, 7, 100, 2, 2, 17, 18,
	7, 111, 2, 2, 18, 19, 7, 119, 2, 2, 19, 24, 7, 110, 2, 2, 20, 21, 7, 108,
	2, 2, 21, 22, 7, 112, 2, 2, 22, 24, 7, 124, 2, 2, 23, 11, 3, 2, 2, 2, 23,
	14, 3, 2, 2, 2, 23, 17, 3, 2, 2, 2, 23, 20, 3, 2, 2, 2, 24, 4, 3, 2, 2,
	2, 25, 26, 4, 99, 106, 2, 26, 6, 3, 2, 2, 2, 27, 29, 7, 47, 2, 2, 28, 27,
	3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 29, 31, 3, 2, 2, 2, 30, 32, 9, 2, 2, 2,
	31, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3,
	2, 2, 2, 34, 8, 3, 2, 2, 2, 35, 37, 9, 3, 2, 2, 36, 35, 3, 2, 2, 2, 37,
	38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 40, 3, 2, 2,
	2, 40, 41, 8, 5, 2, 2, 41, 10, 3, 2, 2, 2, 7, 2, 23, 28, 33, 38, 3, 8,
	2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames []string

var lexerSymbolicNames = []string{
	"", "Opcode", "Reg", "Const", "WS",
}

var lexerRuleNames = []string{
	"Opcode", "Reg", "Const", "WS",
}

type CoAsmLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewCoAsmLexer(input antlr.CharStream) *CoAsmLexer {

	l := new(CoAsmLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "CoAsm.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CoAsmLexer tokens.
const (
	CoAsmLexerOpcode = 1
	CoAsmLexerReg    = 2
	CoAsmLexerConst  = 3
	CoAsmLexerWS     = 4
)
