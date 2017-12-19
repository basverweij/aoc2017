// Code generated from .\SndAsm.g4 by ANTLR 4.7.1. DO NOT EDIT.

package sndasm // SndAsm
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 55, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 7, 2, 24, 10, 2,
	12, 2, 14, 2, 27, 11, 2, 3, 3, 3, 3, 3, 3, 5, 3, 32, 10, 3, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 2, 2, 12, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 2, 5, 3, 2, 3, 4, 3, 2, 5, 8, 3, 2, 10, 11,
	2, 47, 2, 25, 3, 2, 2, 2, 4, 31, 3, 2, 2, 2, 6, 33, 3, 2, 2, 2, 8, 36,
	3, 2, 2, 2, 10, 38, 3, 2, 2, 2, 12, 42, 3, 2, 2, 2, 14, 44, 3, 2, 2, 2,
	16, 48, 3, 2, 2, 2, 18, 50, 3, 2, 2, 2, 20, 52, 3, 2, 2, 2, 22, 24, 5,
	4, 3, 2, 23, 22, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25,
	26, 3, 2, 2, 2, 26, 3, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 32, 5, 6, 4,
	2, 29, 32, 5, 10, 6, 2, 30, 32, 5, 14, 8, 2, 31, 28, 3, 2, 2, 2, 31, 29,
	3, 2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 5, 3, 2, 2, 2, 33, 34, 5, 8, 5, 2,
	34, 35, 5, 20, 11, 2, 35, 7, 3, 2, 2, 2, 36, 37, 9, 2, 2, 2, 37, 9, 3,
	2, 2, 2, 38, 39, 5, 12, 7, 2, 39, 40, 5, 18, 10, 2, 40, 41, 5, 20, 11,
	2, 41, 11, 3, 2, 2, 2, 42, 43, 9, 3, 2, 2, 43, 13, 3, 2, 2, 2, 44, 45,
	5, 16, 9, 2, 45, 46, 5, 20, 11, 2, 46, 47, 5, 20, 11, 2, 47, 15, 3, 2,
	2, 2, 48, 49, 7, 9, 2, 2, 49, 17, 3, 2, 2, 2, 50, 51, 7, 10, 2, 2, 51,
	19, 3, 2, 2, 2, 52, 53, 9, 4, 2, 2, 53, 21, 3, 2, 2, 2, 4, 25, 31,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'snd'", "'rcv'", "'set'", "'add'", "'mul'", "'mod'", "'jgz'",
}
var symbolicNames = []string{
	"", "Snd", "Rcv", "Set", "Add", "Mul", "Mod", "Jgz", "Reg", "Value", "WS",
}

var ruleNames = []string{
	"code", "instruction", "unaryExpression", "unaryInstruction", "binaryRegExpression",
	"binaryRegInstruction", "binaryExpression", "binaryInstruction", "reg",
	"regOrValue",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SndAsmParser struct {
	*antlr.BaseParser
}

func NewSndAsmParser(input antlr.TokenStream) *SndAsmParser {
	this := new(SndAsmParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SndAsm.g4"

	return this
}

// SndAsmParser tokens.
const (
	SndAsmParserEOF   = antlr.TokenEOF
	SndAsmParserSnd   = 1
	SndAsmParserRcv   = 2
	SndAsmParserSet   = 3
	SndAsmParserAdd   = 4
	SndAsmParserMul   = 5
	SndAsmParserMod   = 6
	SndAsmParserJgz   = 7
	SndAsmParserReg   = 8
	SndAsmParserValue = 9
	SndAsmParserWS    = 10
)

// SndAsmParser rules.
const (
	SndAsmParserRULE_code                 = 0
	SndAsmParserRULE_instruction          = 1
	SndAsmParserRULE_unaryExpression      = 2
	SndAsmParserRULE_unaryInstruction     = 3
	SndAsmParserRULE_binaryRegExpression  = 4
	SndAsmParserRULE_binaryRegInstruction = 5
	SndAsmParserRULE_binaryExpression     = 6
	SndAsmParserRULE_binaryInstruction    = 7
	SndAsmParserRULE_reg                  = 8
	SndAsmParserRULE_regOrValue           = 9
)

// ICodeContext is an interface to support dynamic dispatch.
type ICodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCodeContext differentiates from other interfaces.
	IsCodeContext()
}

type CodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCodeContext() *CodeContext {
	var p = new(CodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_code
	return p
}

func (*CodeContext) IsCodeContext() {}

func NewCodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeContext {
	var p = new(CodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_code

	return p
}

func (s *CodeContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *CodeContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *CodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterCode(s)
	}
}

func (s *CodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitCode(s)
	}
}

func (p *SndAsmParser) Code() (localctx ICodeContext) {
	localctx = NewCodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SndAsmParserRULE_code)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SndAsmParserSnd)|(1<<SndAsmParserRcv)|(1<<SndAsmParserSet)|(1<<SndAsmParserAdd)|(1<<SndAsmParserMul)|(1<<SndAsmParserMod)|(1<<SndAsmParserJgz))) != 0 {
		{
			p.SetState(20)
			p.Instruction()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) UnaryExpression() IUnaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *InstructionContext) BinaryRegExpression() IBinaryRegExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryRegExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryRegExpressionContext)
}

func (s *InstructionContext) BinaryExpression() IBinaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryExpressionContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *SndAsmParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SndAsmParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SndAsmParserSnd, SndAsmParserRcv:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.UnaryExpression()
		}

	case SndAsmParserSet, SndAsmParserAdd, SndAsmParserMul, SndAsmParserMod:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.BinaryRegExpression()
		}

	case SndAsmParserJgz:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(28)
			p.BinaryExpression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_unaryExpression
	return p
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) UnaryInstruction() IUnaryInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryInstructionContext)
}

func (s *UnaryExpressionContext) RegOrValue() IRegOrValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegOrValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegOrValueContext)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (p *SndAsmParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SndAsmParserRULE_unaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.UnaryInstruction()
	}
	{
		p.SetState(32)
		p.RegOrValue()
	}

	return localctx
}

// IUnaryInstructionContext is an interface to support dynamic dispatch.
type IUnaryInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryInstructionContext differentiates from other interfaces.
	IsUnaryInstructionContext()
}

type UnaryInstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryInstructionContext() *UnaryInstructionContext {
	var p = new(UnaryInstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_unaryInstruction
	return p
}

func (*UnaryInstructionContext) IsUnaryInstructionContext() {}

func NewUnaryInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryInstructionContext {
	var p = new(UnaryInstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_unaryInstruction

	return p
}

func (s *UnaryInstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryInstructionContext) Snd() antlr.TerminalNode {
	return s.GetToken(SndAsmParserSnd, 0)
}

func (s *UnaryInstructionContext) Rcv() antlr.TerminalNode {
	return s.GetToken(SndAsmParserRcv, 0)
}

func (s *UnaryInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryInstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterUnaryInstruction(s)
	}
}

func (s *UnaryInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitUnaryInstruction(s)
	}
}

func (p *SndAsmParser) UnaryInstruction() (localctx IUnaryInstructionContext) {
	localctx = NewUnaryInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SndAsmParserRULE_unaryInstruction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SndAsmParserSnd || _la == SndAsmParserRcv) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryRegExpressionContext is an interface to support dynamic dispatch.
type IBinaryRegExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryRegExpressionContext differentiates from other interfaces.
	IsBinaryRegExpressionContext()
}

type BinaryRegExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryRegExpressionContext() *BinaryRegExpressionContext {
	var p = new(BinaryRegExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_binaryRegExpression
	return p
}

func (*BinaryRegExpressionContext) IsBinaryRegExpressionContext() {}

func NewBinaryRegExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryRegExpressionContext {
	var p = new(BinaryRegExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_binaryRegExpression

	return p
}

func (s *BinaryRegExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryRegExpressionContext) BinaryRegInstruction() IBinaryRegInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryRegInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryRegInstructionContext)
}

func (s *BinaryRegExpressionContext) Reg() IRegContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegContext)
}

func (s *BinaryRegExpressionContext) RegOrValue() IRegOrValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegOrValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegOrValueContext)
}

func (s *BinaryRegExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryRegExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryRegExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterBinaryRegExpression(s)
	}
}

func (s *BinaryRegExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitBinaryRegExpression(s)
	}
}

func (p *SndAsmParser) BinaryRegExpression() (localctx IBinaryRegExpressionContext) {
	localctx = NewBinaryRegExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SndAsmParserRULE_binaryRegExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(36)
		p.BinaryRegInstruction()
	}
	{
		p.SetState(37)
		p.Reg()
	}
	{
		p.SetState(38)
		p.RegOrValue()
	}

	return localctx
}

// IBinaryRegInstructionContext is an interface to support dynamic dispatch.
type IBinaryRegInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryRegInstructionContext differentiates from other interfaces.
	IsBinaryRegInstructionContext()
}

type BinaryRegInstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryRegInstructionContext() *BinaryRegInstructionContext {
	var p = new(BinaryRegInstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_binaryRegInstruction
	return p
}

func (*BinaryRegInstructionContext) IsBinaryRegInstructionContext() {}

func NewBinaryRegInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryRegInstructionContext {
	var p = new(BinaryRegInstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_binaryRegInstruction

	return p
}

func (s *BinaryRegInstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryRegInstructionContext) Set() antlr.TerminalNode {
	return s.GetToken(SndAsmParserSet, 0)
}

func (s *BinaryRegInstructionContext) Add() antlr.TerminalNode {
	return s.GetToken(SndAsmParserAdd, 0)
}

func (s *BinaryRegInstructionContext) Mul() antlr.TerminalNode {
	return s.GetToken(SndAsmParserMul, 0)
}

func (s *BinaryRegInstructionContext) Mod() antlr.TerminalNode {
	return s.GetToken(SndAsmParserMod, 0)
}

func (s *BinaryRegInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryRegInstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryRegInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterBinaryRegInstruction(s)
	}
}

func (s *BinaryRegInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitBinaryRegInstruction(s)
	}
}

func (p *SndAsmParser) BinaryRegInstruction() (localctx IBinaryRegInstructionContext) {
	localctx = NewBinaryRegInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SndAsmParserRULE_binaryRegInstruction)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(40)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SndAsmParserSet)|(1<<SndAsmParserAdd)|(1<<SndAsmParserMul)|(1<<SndAsmParserMod))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryExpressionContext is an interface to support dynamic dispatch.
type IBinaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryExpressionContext differentiates from other interfaces.
	IsBinaryExpressionContext()
}

type BinaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryExpressionContext() *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_binaryExpression
	return p
}

func (*BinaryExpressionContext) IsBinaryExpressionContext() {}

func NewBinaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_binaryExpression

	return p
}

func (s *BinaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryExpressionContext) BinaryInstruction() IBinaryInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryInstructionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryInstructionContext)
}

func (s *BinaryExpressionContext) AllRegOrValue() []IRegOrValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegOrValueContext)(nil)).Elem())
	var tst = make([]IRegOrValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegOrValueContext)
		}
	}

	return tst
}

func (s *BinaryExpressionContext) RegOrValue(i int) IRegOrValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegOrValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegOrValueContext)
}

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

func (p *SndAsmParser) BinaryExpression() (localctx IBinaryExpressionContext) {
	localctx = NewBinaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SndAsmParserRULE_binaryExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.BinaryInstruction()
	}
	{
		p.SetState(43)
		p.RegOrValue()
	}
	{
		p.SetState(44)
		p.RegOrValue()
	}

	return localctx
}

// IBinaryInstructionContext is an interface to support dynamic dispatch.
type IBinaryInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryInstructionContext differentiates from other interfaces.
	IsBinaryInstructionContext()
}

type BinaryInstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryInstructionContext() *BinaryInstructionContext {
	var p = new(BinaryInstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_binaryInstruction
	return p
}

func (*BinaryInstructionContext) IsBinaryInstructionContext() {}

func NewBinaryInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryInstructionContext {
	var p = new(BinaryInstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_binaryInstruction

	return p
}

func (s *BinaryInstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryInstructionContext) Jgz() antlr.TerminalNode {
	return s.GetToken(SndAsmParserJgz, 0)
}

func (s *BinaryInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryInstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterBinaryInstruction(s)
	}
}

func (s *BinaryInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitBinaryInstruction(s)
	}
}

func (p *SndAsmParser) BinaryInstruction() (localctx IBinaryInstructionContext) {
	localctx = NewBinaryInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SndAsmParserRULE_binaryInstruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(SndAsmParserJgz)
	}

	return localctx
}

// IRegContext is an interface to support dynamic dispatch.
type IRegContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegContext differentiates from other interfaces.
	IsRegContext()
}

type RegContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegContext() *RegContext {
	var p = new(RegContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_reg
	return p
}

func (*RegContext) IsRegContext() {}

func NewRegContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegContext {
	var p = new(RegContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_reg

	return p
}

func (s *RegContext) GetParser() antlr.Parser { return s.parser }

func (s *RegContext) Reg() antlr.TerminalNode {
	return s.GetToken(SndAsmParserReg, 0)
}

func (s *RegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterReg(s)
	}
}

func (s *RegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitReg(s)
	}
}

func (p *SndAsmParser) Reg() (localctx IRegContext) {
	localctx = NewRegContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SndAsmParserRULE_reg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(SndAsmParserReg)
	}

	return localctx
}

// IRegOrValueContext is an interface to support dynamic dispatch.
type IRegOrValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegOrValueContext differentiates from other interfaces.
	IsRegOrValueContext()
}

type RegOrValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegOrValueContext() *RegOrValueContext {
	var p = new(RegOrValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SndAsmParserRULE_regOrValue
	return p
}

func (*RegOrValueContext) IsRegOrValueContext() {}

func NewRegOrValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegOrValueContext {
	var p = new(RegOrValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SndAsmParserRULE_regOrValue

	return p
}

func (s *RegOrValueContext) GetParser() antlr.Parser { return s.parser }

func (s *RegOrValueContext) Reg() antlr.TerminalNode {
	return s.GetToken(SndAsmParserReg, 0)
}

func (s *RegOrValueContext) Value() antlr.TerminalNode {
	return s.GetToken(SndAsmParserValue, 0)
}

func (s *RegOrValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegOrValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegOrValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.EnterRegOrValue(s)
	}
}

func (s *RegOrValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SndAsmListener); ok {
		listenerT.ExitRegOrValue(s)
	}
}

func (p *SndAsmParser) RegOrValue() (localctx IRegOrValueContext) {
	localctx = NewRegOrValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SndAsmParserRULE_regOrValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SndAsmParserReg || _la == SndAsmParserValue) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
