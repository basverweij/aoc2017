// Code generated from .\Turing.g4 by ANTLR 4.7.1. DO NOT EDIT.

package turing

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 16, 222,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 13, 6, 13, 201, 10, 13, 13, 13, 14, 13, 202, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 214, 10, 14,
	3, 15, 6, 15, 217, 10, 15, 13, 15, 14, 15, 218, 3, 15, 3, 15, 2, 2, 16,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23,
	13, 25, 14, 27, 15, 29, 16, 3, 2, 5, 3, 2, 67, 72, 3, 2, 50, 59, 5, 2,
	11, 12, 15, 15, 34, 34, 2, 224, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2,
	2, 2, 3, 31, 3, 2, 2, 2, 5, 47, 3, 2, 2, 2, 7, 49, 3, 2, 2, 2, 9, 86, 3,
	2, 2, 2, 11, 94, 3, 2, 2, 2, 13, 104, 3, 2, 2, 2, 15, 106, 3, 2, 2, 2,
	17, 131, 3, 2, 2, 2, 19, 150, 3, 2, 2, 2, 21, 174, 3, 2, 2, 2, 23, 197,
	3, 2, 2, 2, 25, 200, 3, 2, 2, 2, 27, 213, 3, 2, 2, 2, 29, 216, 3, 2, 2,
	2, 31, 32, 7, 68, 2, 2, 32, 33, 7, 103, 2, 2, 33, 34, 7, 105, 2, 2, 34,
	35, 7, 107, 2, 2, 35, 36, 7, 112, 2, 2, 36, 37, 7, 34, 2, 2, 37, 38, 7,
	107, 2, 2, 38, 39, 7, 112, 2, 2, 39, 40, 7, 34, 2, 2, 40, 41, 7, 117, 2,
	2, 41, 42, 7, 118, 2, 2, 42, 43, 7, 99, 2, 2, 43, 44, 7, 118, 2, 2, 44,
	45, 7, 103, 2, 2, 45, 46, 7, 34, 2, 2, 46, 4, 3, 2, 2, 2, 47, 48, 7, 48,
	2, 2, 48, 6, 3, 2, 2, 2, 49, 50, 7, 82, 2, 2, 50, 51, 7, 103, 2, 2, 51,
	52, 7, 116, 2, 2, 52, 53, 7, 104, 2, 2, 53, 54, 7, 113, 2, 2, 54, 55, 7,
	116, 2, 2, 55, 56, 7, 111, 2, 2, 56, 57, 7, 34, 2, 2, 57, 58, 7, 99, 2,
	2, 58, 59, 7, 34, 2, 2, 59, 60, 7, 102, 2, 2, 60, 61, 7, 107, 2, 2, 61,
	62, 7, 99, 2, 2, 62, 63, 7, 105, 2, 2, 63, 64, 7, 112, 2, 2, 64, 65, 7,
	113, 2, 2, 65, 66, 7, 117, 2, 2, 66, 67, 7, 118, 2, 2, 67, 68, 7, 107,
	2, 2, 68, 69, 7, 101, 2, 2, 69, 70, 7, 34, 2, 2, 70, 71, 7, 101, 2, 2,
	71, 72, 7, 106, 2, 2, 72, 73, 7, 103, 2, 2, 73, 74, 7, 101, 2, 2, 74, 75,
	7, 109, 2, 2, 75, 76, 7, 117, 2, 2, 76, 77, 7, 119, 2, 2, 77, 78, 7, 111,
	2, 2, 78, 79, 7, 34, 2, 2, 79, 80, 7, 99, 2, 2, 80, 81, 7, 104, 2, 2, 81,
	82, 7, 118, 2, 2, 82, 83, 7, 103, 2, 2, 83, 84, 7, 116, 2, 2, 84, 85, 7,
	34, 2, 2, 85, 8, 3, 2, 2, 2, 86, 87, 7, 34, 2, 2, 87, 88, 7, 117, 2, 2,
	88, 89, 7, 118, 2, 2, 89, 90, 7, 103, 2, 2, 90, 91, 7, 114, 2, 2, 91, 92,
	7, 117, 2, 2, 92, 93, 7, 48, 2, 2, 93, 10, 3, 2, 2, 2, 94, 95, 7, 75, 2,
	2, 95, 96, 7, 112, 2, 2, 96, 97, 7, 34, 2, 2, 97, 98, 7, 117, 2, 2, 98,
	99, 7, 118, 2, 2, 99, 100, 7, 99, 2, 2, 100, 101, 7, 118, 2, 2, 101, 102,
	7, 103, 2, 2, 102, 103, 7, 34, 2, 2, 103, 12, 3, 2, 2, 2, 104, 105, 7,
	60, 2, 2, 105, 14, 3, 2, 2, 2, 106, 107, 7, 75, 2, 2, 107, 108, 7, 104,
	2, 2, 108, 109, 7, 34, 2, 2, 109, 110, 7, 118, 2, 2, 110, 111, 7, 106,
	2, 2, 111, 112, 7, 103, 2, 2, 112, 113, 7, 34, 2, 2, 113, 114, 7, 101,
	2, 2, 114, 115, 7, 119, 2, 2, 115, 116, 7, 116, 2, 2, 116, 117, 7, 116,
	2, 2, 117, 118, 7, 103, 2, 2, 118, 119, 7, 112, 2, 2, 119, 120, 7, 118,
	2, 2, 120, 121, 7, 34, 2, 2, 121, 122, 7, 120, 2, 2, 122, 123, 7, 99, 2,
	2, 123, 124, 7, 110, 2, 2, 124, 125, 7, 119, 2, 2, 125, 126, 7, 103, 2,
	2, 126, 127, 7, 34, 2, 2, 127, 128, 7, 107, 2, 2, 128, 129, 7, 117, 2,
	2, 129, 130, 7, 34, 2, 2, 130, 16, 3, 2, 2, 2, 131, 132, 7, 47, 2, 2, 132,
	133, 7, 34, 2, 2, 133, 134, 7, 89, 2, 2, 134, 135, 7, 116, 2, 2, 135, 136,
	7, 107, 2, 2, 136, 137, 7, 118, 2, 2, 137, 138, 7, 103, 2, 2, 138, 139,
	7, 34, 2, 2, 139, 140, 7, 118, 2, 2, 140, 141, 7, 106, 2, 2, 141, 142,
	7, 103, 2, 2, 142, 143, 7, 34, 2, 2, 143, 144, 7, 120, 2, 2, 144, 145,
	7, 99, 2, 2, 145, 146, 7, 110, 2, 2, 146, 147, 7, 119, 2, 2, 147, 148,
	7, 103, 2, 2, 148, 149, 7, 34, 2, 2, 149, 18, 3, 2, 2, 2, 150, 151, 7,
	47, 2, 2, 151, 152, 7, 34, 2, 2, 152, 153, 7, 79, 2, 2, 153, 154, 7, 113,
	2, 2, 154, 155, 7, 120, 2, 2, 155, 156, 7, 103, 2, 2, 156, 157, 7, 34,
	2, 2, 157, 158, 7, 113, 2, 2, 158, 159, 7, 112, 2, 2, 159, 160, 7, 103,
	2, 2, 160, 161, 7, 34, 2, 2, 161, 162, 7, 117, 2, 2, 162, 163, 7, 110,
	2, 2, 163, 164, 7, 113, 2, 2, 164, 165, 7, 118, 2, 2, 165, 166, 7, 34,
	2, 2, 166, 167, 7, 118, 2, 2, 167, 168, 7, 113, 2, 2, 168, 169, 7, 34,
	2, 2, 169, 170, 7, 118, 2, 2, 170, 171, 7, 106, 2, 2, 171, 172, 7, 103,
	2, 2, 172, 173, 7, 34, 2, 2, 173, 20, 3, 2, 2, 2, 174, 175, 7, 47, 2, 2,
	175, 176, 7, 34, 2, 2, 176, 177, 7, 69, 2, 2, 177, 178, 7, 113, 2, 2, 178,
	179, 7, 112, 2, 2, 179, 180, 7, 118, 2, 2, 180, 181, 7, 107, 2, 2, 181,
	182, 7, 112, 2, 2, 182, 183, 7, 119, 2, 2, 183, 184, 7, 103, 2, 2, 184,
	185, 7, 34, 2, 2, 185, 186, 7, 121, 2, 2, 186, 187, 7, 107, 2, 2, 187,
	188, 7, 118, 2, 2, 188, 189, 7, 106, 2, 2, 189, 190, 7, 34, 2, 2, 190,
	191, 7, 117, 2, 2, 191, 192, 7, 118, 2, 2, 192, 193, 7, 99, 2, 2, 193,
	194, 7, 118, 2, 2, 194, 195, 7, 103, 2, 2, 195, 196, 7, 34, 2, 2, 196,
	22, 3, 2, 2, 2, 197, 198, 9, 2, 2, 2, 198, 24, 3, 2, 2, 2, 199, 201, 9,
	3, 2, 2, 200, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 200, 3, 2, 2,
	2, 202, 203, 3, 2, 2, 2, 203, 26, 3, 2, 2, 2, 204, 205, 7, 110, 2, 2, 205,
	206, 7, 103, 2, 2, 206, 207, 7, 104, 2, 2, 207, 214, 7, 118, 2, 2, 208,
	209, 7, 116, 2, 2, 209, 210, 7, 107, 2, 2, 210, 211, 7, 105, 2, 2, 211,
	212, 7, 106, 2, 2, 212, 214, 7, 118, 2, 2, 213, 204, 3, 2, 2, 2, 213, 208,
	3, 2, 2, 2, 214, 28, 3, 2, 2, 2, 215, 217, 9, 4, 2, 2, 216, 215, 3, 2,
	2, 2, 217, 218, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2,
	219, 220, 3, 2, 2, 2, 220, 221, 8, 15, 2, 2, 221, 30, 3, 2, 2, 2, 6, 2,
	202, 213, 218, 3, 8, 2, 2,
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
	"", "'Begin in state '", "'.'", "'Perform a diagnostic checksum after '",
	"' steps.'", "'In state '", "':'", "'If the current value is '", "'- Write the value '",
	"'- Move one slot to the '", "'- Continue with state '",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "State", "Num", "Dir", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "State", "Num", "Dir", "WS",
}

type TuringLexer struct {
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

func NewTuringLexer(input antlr.CharStream) *TuringLexer {

	l := new(TuringLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Turing.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TuringLexer tokens.
const (
	TuringLexerT__0  = 1
	TuringLexerT__1  = 2
	TuringLexerT__2  = 3
	TuringLexerT__3  = 4
	TuringLexerT__4  = 5
	TuringLexerT__5  = 6
	TuringLexerT__6  = 7
	TuringLexerT__7  = 8
	TuringLexerT__8  = 9
	TuringLexerT__9  = 10
	TuringLexerState = 11
	TuringLexerNum   = 12
	TuringLexerDir   = 13
	TuringLexerWS    = 14
)