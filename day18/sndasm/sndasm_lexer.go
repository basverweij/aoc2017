// Code generated from .\SndAsm.g4 by ANTLR 4.7.1. DO NOT EDIT.

package sndasm

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 68, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 5, 10, 55, 10, 10, 3, 10, 6, 10, 58, 10, 10, 13,
	10, 14, 10, 59, 3, 11, 6, 11, 63, 10, 11, 13, 11, 14, 11, 64, 3, 11, 3,
	11, 2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 3, 2, 5, 3, 2, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 70, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2,
	2, 5, 27, 3, 2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 35, 3, 2, 2, 2, 11, 39, 3,
	2, 2, 2, 13, 43, 3, 2, 2, 2, 15, 47, 3, 2, 2, 2, 17, 51, 3, 2, 2, 2, 19,
	54, 3, 2, 2, 2, 21, 62, 3, 2, 2, 2, 23, 24, 7, 117, 2, 2, 24, 25, 7, 112,
	2, 2, 25, 26, 7, 102, 2, 2, 26, 4, 3, 2, 2, 2, 27, 28, 7, 116, 2, 2, 28,
	29, 7, 101, 2, 2, 29, 30, 7, 120, 2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 117,
	2, 2, 32, 33, 7, 103, 2, 2, 33, 34, 7, 118, 2, 2, 34, 8, 3, 2, 2, 2, 35,
	36, 7, 99, 2, 2, 36, 37, 7, 102, 2, 2, 37, 38, 7, 102, 2, 2, 38, 10, 3,
	2, 2, 2, 39, 40, 7, 111, 2, 2, 40, 41, 7, 119, 2, 2, 41, 42, 7, 110, 2,
	2, 42, 12, 3, 2, 2, 2, 43, 44, 7, 111, 2, 2, 44, 45, 7, 113, 2, 2, 45,
	46, 7, 102, 2, 2, 46, 14, 3, 2, 2, 2, 47, 48, 7, 108, 2, 2, 48, 49, 7,
	105, 2, 2, 49, 50, 7, 124, 2, 2, 50, 16, 3, 2, 2, 2, 51, 52, 9, 2, 2, 2,
	52, 18, 3, 2, 2, 2, 53, 55, 7, 47, 2, 2, 54, 53, 3, 2, 2, 2, 54, 55, 3,
	2, 2, 2, 55, 57, 3, 2, 2, 2, 56, 58, 9, 3, 2, 2, 57, 56, 3, 2, 2, 2, 58,
	59, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 20, 3, 2, 2,
	2, 61, 63, 9, 4, 2, 2, 62, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 62,
	3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 67, 8, 11, 2, 2,
	67, 22, 3, 2, 2, 2, 6, 2, 54, 59, 64, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'snd'", "'rcv'", "'set'", "'add'", "'mul'", "'mod'", "'jgz'",
}

var lexerSymbolicNames = []string{
	"", "Snd", "Rcv", "Set", "Add", "Mul", "Mod", "Jgz", "Reg", "Value", "WS",
}

var lexerRuleNames = []string{
	"Snd", "Rcv", "Set", "Add", "Mul", "Mod", "Jgz", "Reg", "Value", "WS",
}

type SndAsmLexer struct {
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

func NewSndAsmLexer(input antlr.CharStream) *SndAsmLexer {

	l := new(SndAsmLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SndAsm.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SndAsmLexer tokens.
const (
	SndAsmLexerSnd   = 1
	SndAsmLexerRcv   = 2
	SndAsmLexerSet   = 3
	SndAsmLexerAdd   = 4
	SndAsmLexerMul   = 5
	SndAsmLexerMod   = 6
	SndAsmLexerJgz   = 7
	SndAsmLexerReg   = 8
	SndAsmLexerValue = 9
	SndAsmLexerWS    = 10
)
