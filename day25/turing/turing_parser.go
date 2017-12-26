// Code generated from .\Turing.g4 by ANTLR 4.7.1. DO NOT EDIT.

package turing // Turing
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 67, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 7, 2, 24, 10, 2, 12,
	2, 14, 2, 27, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 5, 3, 5, 7, 5, 41, 10, 5, 12, 5, 14, 5, 44, 11, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 2, 2, 11, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 2, 2, 2, 59, 2, 20, 3, 2, 2, 2, 4, 28, 3, 2,
	2, 2, 6, 32, 3, 2, 2, 2, 8, 36, 3, 2, 2, 2, 10, 45, 3, 2, 2, 2, 12, 50,
	3, 2, 2, 2, 14, 54, 3, 2, 2, 2, 16, 58, 3, 2, 2, 2, 18, 62, 3, 2, 2, 2,
	20, 21, 5, 4, 3, 2, 21, 25, 5, 6, 4, 2, 22, 24, 5, 8, 5, 2, 23, 22, 3,
	2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26,
	3, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 29, 7, 3, 2, 2, 29, 30, 7, 13, 2,
	2, 30, 31, 7, 4, 2, 2, 31, 5, 3, 2, 2, 2, 32, 33, 7, 5, 2, 2, 33, 34, 7,
	14, 2, 2, 34, 35, 7, 6, 2, 2, 35, 7, 3, 2, 2, 2, 36, 37, 7, 7, 2, 2, 37,
	38, 7, 13, 2, 2, 38, 42, 7, 8, 2, 2, 39, 41, 5, 10, 6, 2, 40, 39, 3, 2,
	2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 9,
	3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 46, 7, 9, 2, 2, 46, 47, 7, 14, 2, 2,
	47, 48, 7, 8, 2, 2, 48, 49, 5, 12, 7, 2, 49, 11, 3, 2, 2, 2, 50, 51, 5,
	14, 8, 2, 51, 52, 5, 16, 9, 2, 52, 53, 5, 18, 10, 2, 53, 13, 3, 2, 2, 2,
	54, 55, 7, 10, 2, 2, 55, 56, 7, 14, 2, 2, 56, 57, 7, 4, 2, 2, 57, 15, 3,
	2, 2, 2, 58, 59, 7, 11, 2, 2, 59, 60, 7, 15, 2, 2, 60, 61, 7, 4, 2, 2,
	61, 17, 3, 2, 2, 2, 62, 63, 7, 12, 2, 2, 63, 64, 7, 13, 2, 2, 64, 65, 7,
	4, 2, 2, 65, 19, 3, 2, 2, 2, 4, 25, 42,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'Begin in state '", "'.'", "'Perform a diagnostic checksum after '",
	"' steps.'", "'In state '", "':'", "'If the current value is '", "'- Write the value '",
	"'- Move one slot to the '", "'- Continue with state '",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "State", "Num", "Dir", "WS",
}

var ruleNames = []string{
	"blueprint", "startState", "diagnostics", "state", "stateRules", "rules",
	"write", "move", "nextState",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type TuringParser struct {
	*antlr.BaseParser
}

func NewTuringParser(input antlr.TokenStream) *TuringParser {
	this := new(TuringParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Turing.g4"

	return this
}

// TuringParser tokens.
const (
	TuringParserEOF   = antlr.TokenEOF
	TuringParserT__0  = 1
	TuringParserT__1  = 2
	TuringParserT__2  = 3
	TuringParserT__3  = 4
	TuringParserT__4  = 5
	TuringParserT__5  = 6
	TuringParserT__6  = 7
	TuringParserT__7  = 8
	TuringParserT__8  = 9
	TuringParserT__9  = 10
	TuringParserState = 11
	TuringParserNum   = 12
	TuringParserDir   = 13
	TuringParserWS    = 14
)

// TuringParser rules.
const (
	TuringParserRULE_blueprint   = 0
	TuringParserRULE_startState  = 1
	TuringParserRULE_diagnostics = 2
	TuringParserRULE_state       = 3
	TuringParserRULE_stateRules  = 4
	TuringParserRULE_rules       = 5
	TuringParserRULE_write       = 6
	TuringParserRULE_move        = 7
	TuringParserRULE_nextState   = 8
)

// IBlueprintContext is an interface to support dynamic dispatch.
type IBlueprintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlueprintContext differentiates from other interfaces.
	IsBlueprintContext()
}

type BlueprintContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlueprintContext() *BlueprintContext {
	var p = new(BlueprintContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_blueprint
	return p
}

func (*BlueprintContext) IsBlueprintContext() {}

func NewBlueprintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlueprintContext {
	var p = new(BlueprintContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_blueprint

	return p
}

func (s *BlueprintContext) GetParser() antlr.Parser { return s.parser }

func (s *BlueprintContext) StartState() IStartStateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartStateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStartStateContext)
}

func (s *BlueprintContext) Diagnostics() IDiagnosticsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDiagnosticsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDiagnosticsContext)
}

func (s *BlueprintContext) AllState() []IStateContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStateContext)(nil)).Elem())
	var tst = make([]IStateContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStateContext)
		}
	}

	return tst
}

func (s *BlueprintContext) State(i int) IStateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStateContext)
}

func (s *BlueprintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlueprintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlueprintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterBlueprint(s)
	}
}

func (s *BlueprintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitBlueprint(s)
	}
}

func (p *TuringParser) Blueprint() (localctx IBlueprintContext) {
	localctx = NewBlueprintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TuringParserRULE_blueprint)
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
		p.SetState(18)
		p.StartState()
	}
	{
		p.SetState(19)
		p.Diagnostics()
	}
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TuringParserT__4 {
		{
			p.SetState(20)
			p.State()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStartStateContext is an interface to support dynamic dispatch.
type IStartStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartStateContext differentiates from other interfaces.
	IsStartStateContext()
}

type StartStateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartStateContext() *StartStateContext {
	var p = new(StartStateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_startState
	return p
}

func (*StartStateContext) IsStartStateContext() {}

func NewStartStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartStateContext {
	var p = new(StartStateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_startState

	return p
}

func (s *StartStateContext) GetParser() antlr.Parser { return s.parser }

func (s *StartStateContext) State() antlr.TerminalNode {
	return s.GetToken(TuringParserState, 0)
}

func (s *StartStateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartStateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartStateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterStartState(s)
	}
}

func (s *StartStateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitStartState(s)
	}
}

func (p *TuringParser) StartState() (localctx IStartStateContext) {
	localctx = NewStartStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TuringParserRULE_startState)

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
		p.SetState(26)
		p.Match(TuringParserT__0)
	}
	{
		p.SetState(27)
		p.Match(TuringParserState)
	}
	{
		p.SetState(28)
		p.Match(TuringParserT__1)
	}

	return localctx
}

// IDiagnosticsContext is an interface to support dynamic dispatch.
type IDiagnosticsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDiagnosticsContext differentiates from other interfaces.
	IsDiagnosticsContext()
}

type DiagnosticsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDiagnosticsContext() *DiagnosticsContext {
	var p = new(DiagnosticsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_diagnostics
	return p
}

func (*DiagnosticsContext) IsDiagnosticsContext() {}

func NewDiagnosticsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DiagnosticsContext {
	var p = new(DiagnosticsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_diagnostics

	return p
}

func (s *DiagnosticsContext) GetParser() antlr.Parser { return s.parser }

func (s *DiagnosticsContext) Num() antlr.TerminalNode {
	return s.GetToken(TuringParserNum, 0)
}

func (s *DiagnosticsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DiagnosticsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DiagnosticsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterDiagnostics(s)
	}
}

func (s *DiagnosticsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitDiagnostics(s)
	}
}

func (p *TuringParser) Diagnostics() (localctx IDiagnosticsContext) {
	localctx = NewDiagnosticsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TuringParserRULE_diagnostics)

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
		p.SetState(30)
		p.Match(TuringParserT__2)
	}
	{
		p.SetState(31)
		p.Match(TuringParserNum)
	}
	{
		p.SetState(32)
		p.Match(TuringParserT__3)
	}

	return localctx
}

// IStateContext is an interface to support dynamic dispatch.
type IStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateContext differentiates from other interfaces.
	IsStateContext()
}

type StateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateContext() *StateContext {
	var p = new(StateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_state
	return p
}

func (*StateContext) IsStateContext() {}

func NewStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateContext {
	var p = new(StateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_state

	return p
}

func (s *StateContext) GetParser() antlr.Parser { return s.parser }

func (s *StateContext) State() antlr.TerminalNode {
	return s.GetToken(TuringParserState, 0)
}

func (s *StateContext) AllStateRules() []IStateRulesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStateRulesContext)(nil)).Elem())
	var tst = make([]IStateRulesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStateRulesContext)
		}
	}

	return tst
}

func (s *StateContext) StateRules(i int) IStateRulesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateRulesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStateRulesContext)
}

func (s *StateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterState(s)
	}
}

func (s *StateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitState(s)
	}
}

func (p *TuringParser) State() (localctx IStateContext) {
	localctx = NewStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TuringParserRULE_state)
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
		p.Match(TuringParserT__4)
	}
	{
		p.SetState(35)
		p.Match(TuringParserState)
	}
	{
		p.SetState(36)
		p.Match(TuringParserT__5)
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == TuringParserT__6 {
		{
			p.SetState(37)
			p.StateRules()
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStateRulesContext is an interface to support dynamic dispatch.
type IStateRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateRulesContext differentiates from other interfaces.
	IsStateRulesContext()
}

type StateRulesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateRulesContext() *StateRulesContext {
	var p = new(StateRulesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_stateRules
	return p
}

func (*StateRulesContext) IsStateRulesContext() {}

func NewStateRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateRulesContext {
	var p = new(StateRulesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_stateRules

	return p
}

func (s *StateRulesContext) GetParser() antlr.Parser { return s.parser }

func (s *StateRulesContext) Num() antlr.TerminalNode {
	return s.GetToken(TuringParserNum, 0)
}

func (s *StateRulesContext) Rules() IRulesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRulesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRulesContext)
}

func (s *StateRulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateRulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateRulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterStateRules(s)
	}
}

func (s *StateRulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitStateRules(s)
	}
}

func (p *TuringParser) StateRules() (localctx IStateRulesContext) {
	localctx = NewStateRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TuringParserRULE_stateRules)

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
		p.SetState(43)
		p.Match(TuringParserT__6)
	}
	{
		p.SetState(44)
		p.Match(TuringParserNum)
	}
	{
		p.SetState(45)
		p.Match(TuringParserT__5)
	}
	{
		p.SetState(46)
		p.Rules()
	}

	return localctx
}

// IRulesContext is an interface to support dynamic dispatch.
type IRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulesContext differentiates from other interfaces.
	IsRulesContext()
}

type RulesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesContext() *RulesContext {
	var p = new(RulesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_rules
	return p
}

func (*RulesContext) IsRulesContext() {}

func NewRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesContext {
	var p = new(RulesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_rules

	return p
}

func (s *RulesContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesContext) Write() IWriteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWriteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWriteContext)
}

func (s *RulesContext) Move() IMoveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMoveContext)
}

func (s *RulesContext) NextState() INextStateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INextStateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INextStateContext)
}

func (s *RulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterRules(s)
	}
}

func (s *RulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitRules(s)
	}
}

func (p *TuringParser) Rules() (localctx IRulesContext) {
	localctx = NewRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TuringParserRULE_rules)

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
		p.Write()
	}
	{
		p.SetState(49)
		p.Move()
	}
	{
		p.SetState(50)
		p.NextState()
	}

	return localctx
}

// IWriteContext is an interface to support dynamic dispatch.
type IWriteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWriteContext differentiates from other interfaces.
	IsWriteContext()
}

type WriteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWriteContext() *WriteContext {
	var p = new(WriteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_write
	return p
}

func (*WriteContext) IsWriteContext() {}

func NewWriteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WriteContext {
	var p = new(WriteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_write

	return p
}

func (s *WriteContext) GetParser() antlr.Parser { return s.parser }

func (s *WriteContext) Num() antlr.TerminalNode {
	return s.GetToken(TuringParserNum, 0)
}

func (s *WriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WriteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterWrite(s)
	}
}

func (s *WriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitWrite(s)
	}
}

func (p *TuringParser) Write() (localctx IWriteContext) {
	localctx = NewWriteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TuringParserRULE_write)

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
		p.SetState(52)
		p.Match(TuringParserT__7)
	}
	{
		p.SetState(53)
		p.Match(TuringParserNum)
	}
	{
		p.SetState(54)
		p.Match(TuringParserT__1)
	}

	return localctx
}

// IMoveContext is an interface to support dynamic dispatch.
type IMoveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoveContext differentiates from other interfaces.
	IsMoveContext()
}

type MoveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoveContext() *MoveContext {
	var p = new(MoveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_move
	return p
}

func (*MoveContext) IsMoveContext() {}

func NewMoveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoveContext {
	var p = new(MoveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_move

	return p
}

func (s *MoveContext) GetParser() antlr.Parser { return s.parser }

func (s *MoveContext) Dir() antlr.TerminalNode {
	return s.GetToken(TuringParserDir, 0)
}

func (s *MoveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterMove(s)
	}
}

func (s *MoveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitMove(s)
	}
}

func (p *TuringParser) Move() (localctx IMoveContext) {
	localctx = NewMoveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TuringParserRULE_move)

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
		p.SetState(56)
		p.Match(TuringParserT__8)
	}
	{
		p.SetState(57)
		p.Match(TuringParserDir)
	}
	{
		p.SetState(58)
		p.Match(TuringParserT__1)
	}

	return localctx
}

// INextStateContext is an interface to support dynamic dispatch.
type INextStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNextStateContext differentiates from other interfaces.
	IsNextStateContext()
}

type NextStateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNextStateContext() *NextStateContext {
	var p = new(NextStateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TuringParserRULE_nextState
	return p
}

func (*NextStateContext) IsNextStateContext() {}

func NewNextStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NextStateContext {
	var p = new(NextStateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TuringParserRULE_nextState

	return p
}

func (s *NextStateContext) GetParser() antlr.Parser { return s.parser }

func (s *NextStateContext) State() antlr.TerminalNode {
	return s.GetToken(TuringParserState, 0)
}

func (s *NextStateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NextStateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NextStateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.EnterNextState(s)
	}
}

func (s *NextStateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TuringListener); ok {
		listenerT.ExitNextState(s)
	}
}

func (p *TuringParser) NextState() (localctx INextStateContext) {
	localctx = NewNextStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TuringParserRULE_nextState)

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
		p.SetState(60)
		p.Match(TuringParserT__9)
	}
	{
		p.SetState(61)
		p.Match(TuringParserState)
	}
	{
		p.SetState(62)
		p.Match(TuringParserT__1)
	}

	return localctx
}
