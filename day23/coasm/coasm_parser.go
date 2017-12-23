// Code generated from .\CoAsm.g4 by ANTLR 4.7.1. DO NOT EDIT.

package coasm // CoAsm
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 21, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 7, 2, 10, 10, 2, 12, 2, 14, 2, 13,
	11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 3,
	3, 2, 4, 5, 2, 18, 2, 11, 3, 2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 18, 3, 2, 2,
	2, 8, 10, 5, 4, 3, 2, 9, 8, 3, 2, 2, 2, 10, 13, 3, 2, 2, 2, 11, 9, 3, 2,
	2, 2, 11, 12, 3, 2, 2, 2, 12, 3, 3, 2, 2, 2, 13, 11, 3, 2, 2, 2, 14, 15,
	7, 3, 2, 2, 15, 16, 5, 6, 4, 2, 16, 17, 5, 6, 4, 2, 17, 5, 3, 2, 2, 2,
	18, 19, 9, 2, 2, 2, 19, 7, 3, 2, 2, 2, 3, 11,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames []string

var symbolicNames = []string{
	"", "Opcode", "Reg", "Const", "WS",
}

var ruleNames = []string{
	"program", "instruction", "regOrConst",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type CoAsmParser struct {
	*antlr.BaseParser
}

func NewCoAsmParser(input antlr.TokenStream) *CoAsmParser {
	this := new(CoAsmParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "CoAsm.g4"

	return this
}

// CoAsmParser tokens.
const (
	CoAsmParserEOF    = antlr.TokenEOF
	CoAsmParserOpcode = 1
	CoAsmParserReg    = 2
	CoAsmParserConst  = 3
	CoAsmParserWS     = 4
)

// CoAsmParser rules.
const (
	CoAsmParserRULE_program     = 0
	CoAsmParserRULE_instruction = 1
	CoAsmParserRULE_regOrConst  = 2
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CoAsmParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CoAsmParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *ProgramContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoAsmListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoAsmListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *CoAsmParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CoAsmParserRULE_program)
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
	p.SetState(9)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CoAsmParserOpcode {
		{
			p.SetState(6)
			p.Instruction()
		}

		p.SetState(11)
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
	p.RuleIndex = CoAsmParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CoAsmParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) Opcode() antlr.TerminalNode {
	return s.GetToken(CoAsmParserOpcode, 0)
}

func (s *InstructionContext) AllRegOrConst() []IRegOrConstContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRegOrConstContext)(nil)).Elem())
	var tst = make([]IRegOrConstContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRegOrConstContext)
		}
	}

	return tst
}

func (s *InstructionContext) RegOrConst(i int) IRegOrConstContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegOrConstContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRegOrConstContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoAsmListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoAsmListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *CoAsmParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CoAsmParserRULE_instruction)

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
		p.SetState(12)
		p.Match(CoAsmParserOpcode)
	}
	{
		p.SetState(13)
		p.RegOrConst()
	}
	{
		p.SetState(14)
		p.RegOrConst()
	}

	return localctx
}

// IRegOrConstContext is an interface to support dynamic dispatch.
type IRegOrConstContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegOrConstContext differentiates from other interfaces.
	IsRegOrConstContext()
}

type RegOrConstContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegOrConstContext() *RegOrConstContext {
	var p = new(RegOrConstContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CoAsmParserRULE_regOrConst
	return p
}

func (*RegOrConstContext) IsRegOrConstContext() {}

func NewRegOrConstContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegOrConstContext {
	var p = new(RegOrConstContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CoAsmParserRULE_regOrConst

	return p
}

func (s *RegOrConstContext) GetParser() antlr.Parser { return s.parser }

func (s *RegOrConstContext) Reg() antlr.TerminalNode {
	return s.GetToken(CoAsmParserReg, 0)
}

func (s *RegOrConstContext) Const() antlr.TerminalNode {
	return s.GetToken(CoAsmParserConst, 0)
}

func (s *RegOrConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegOrConstContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegOrConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoAsmListener); ok {
		listenerT.EnterRegOrConst(s)
	}
}

func (s *RegOrConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CoAsmListener); ok {
		listenerT.ExitRegOrConst(s)
	}
}

func (p *CoAsmParser) RegOrConst() (localctx IRegOrConstContext) {
	localctx = NewRegOrConstContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CoAsmParserRULE_regOrConst)
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
		p.SetState(16)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CoAsmParserReg || _la == CoAsmParserConst) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
