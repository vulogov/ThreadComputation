// Code generated from ThreadComputation.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // ThreadComputation

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 65, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 7, 2, 18, 10, 2, 12, 2, 14, 2, 21, 11, 2, 3, 3, 3, 3, 3,
	3, 5, 3, 26, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 32, 10, 4, 12, 4, 14,
	4, 35, 11, 4, 3, 4, 3, 4, 3, 5, 5, 5, 40, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	45, 10, 5, 12, 5, 14, 5, 48, 11, 5, 3, 5, 5, 5, 51, 10, 5, 3, 6, 3, 6,
	3, 6, 5, 6, 56, 10, 6, 3, 7, 3, 7, 5, 7, 60, 10, 7, 3, 8, 3, 8, 3, 8, 3,
	8, 2, 2, 9, 2, 4, 6, 8, 10, 12, 14, 2, 2, 2, 67, 2, 19, 3, 2, 2, 2, 4,
	25, 3, 2, 2, 2, 6, 27, 3, 2, 2, 2, 8, 39, 3, 2, 2, 2, 10, 55, 3, 2, 2,
	2, 12, 59, 3, 2, 2, 2, 14, 61, 3, 2, 2, 2, 16, 18, 5, 4, 3, 2, 17, 16,
	3, 2, 2, 2, 18, 21, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2,
	20, 3, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 22, 26, 5, 8, 5, 2, 23, 26, 5, 6,
	4, 2, 24, 26, 5, 14, 8, 2, 25, 22, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25,
	24, 3, 2, 2, 2, 26, 5, 3, 2, 2, 2, 27, 28, 7, 3, 2, 2, 28, 29, 7, 8, 2,
	2, 29, 33, 7, 4, 2, 2, 30, 32, 5, 10, 6, 2, 31, 30, 3, 2, 2, 2, 32, 35,
	3, 2, 2, 2, 33, 31, 3, 2, 2, 2, 33, 34, 3, 2, 2, 2, 34, 36, 3, 2, 2, 2,
	35, 33, 3, 2, 2, 2, 36, 37, 7, 5, 2, 2, 37, 7, 3, 2, 2, 2, 38, 40, 7, 16,
	2, 2, 39, 38, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 50,
	7, 8, 2, 2, 42, 46, 7, 3, 2, 2, 43, 45, 5, 12, 7, 2, 44, 43, 3, 2, 2, 2,
	45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3,
	2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 51, 7, 6, 2, 2, 50, 42, 3, 2, 2, 2, 50,
	51, 3, 2, 2, 2, 51, 9, 3, 2, 2, 2, 52, 56, 5, 8, 5, 2, 53, 56, 5, 6, 4,
	2, 54, 56, 5, 14, 8, 2, 55, 52, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 54,
	3, 2, 2, 2, 56, 11, 3, 2, 2, 2, 57, 60, 5, 8, 5, 2, 58, 60, 5, 14, 8, 2,
	59, 57, 3, 2, 2, 2, 59, 58, 3, 2, 2, 2, 60, 13, 3, 2, 2, 2, 61, 62, 7,
	7, 2, 2, 62, 63, 7, 8, 2, 2, 63, 15, 3, 2, 2, 2, 10, 19, 25, 33, 39, 46,
	50, 55, 59,
}
var literalNames = []string{
	"", "'['", "':'", "';;'", "']'", "'#'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "FUNC_NAME", "OPS", "NAME", "INTEGER", "DECIMAL_INTEGER",
	"FLOAT_NUMBER", "STRING", "OP", "MOD", "BLOCK_COMMENT", "WhiteSpace", "NewLine",
}

var ruleNames = []string{
	"expressions", "root_term", "ns", "fun", "ns_term", "fun_term", "pos_term",
}

type ThreadComputationParser struct {
	*antlr.BaseParser
}

// NewThreadComputationParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ThreadComputationParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewThreadComputationParser(input antlr.TokenStream) *ThreadComputationParser {
	this := new(ThreadComputationParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ThreadComputation.g4"

	return this
}

// ThreadComputationParser tokens.
const (
	ThreadComputationParserEOF             = antlr.TokenEOF
	ThreadComputationParserT__0            = 1
	ThreadComputationParserT__1            = 2
	ThreadComputationParserT__2            = 3
	ThreadComputationParserT__3            = 4
	ThreadComputationParserT__4            = 5
	ThreadComputationParserFUNC_NAME       = 6
	ThreadComputationParserOPS             = 7
	ThreadComputationParserNAME            = 8
	ThreadComputationParserINTEGER         = 9
	ThreadComputationParserDECIMAL_INTEGER = 10
	ThreadComputationParserFLOAT_NUMBER    = 11
	ThreadComputationParserSTRING          = 12
	ThreadComputationParserOP              = 13
	ThreadComputationParserMOD             = 14
	ThreadComputationParserBLOCK_COMMENT   = 15
	ThreadComputationParserWhiteSpace      = 16
	ThreadComputationParserNewLine         = 17
)

// ThreadComputationParser rules.
const (
	ThreadComputationParserRULE_expressions = 0
	ThreadComputationParserRULE_root_term   = 1
	ThreadComputationParserRULE_ns          = 2
	ThreadComputationParserRULE_fun         = 3
	ThreadComputationParserRULE_ns_term     = 4
	ThreadComputationParserRULE_fun_term    = 5
	ThreadComputationParserRULE_pos_term    = 6
)

// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_expressions

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) AllRoot_term() []IRoot_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRoot_termContext)(nil)).Elem())
	var tst = make([]IRoot_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRoot_termContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Root_term(i int) IRoot_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRoot_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRoot_termContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (p *ThreadComputationParser) Expressions() (localctx IExpressionsContext) {
	this := p
	_ = this

	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ThreadComputationParserRULE_expressions)
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
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__0)|(1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
		{
			p.SetState(14)
			p.Root_term()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRoot_termContext is an interface to support dynamic dispatch.
type IRoot_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoot_termContext differentiates from other interfaces.
	IsRoot_termContext()
}

type Root_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoot_termContext() *Root_termContext {
	var p = new(Root_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_root_term
	return p
}

func (*Root_termContext) IsRoot_termContext() {}

func NewRoot_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Root_termContext {
	var p = new(Root_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_root_term

	return p
}

func (s *Root_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Root_termContext) Fun() IFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *Root_termContext) Ns() INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *Root_termContext) Pos_term() IPos_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPos_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPos_termContext)
}

func (s *Root_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Root_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Root_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterRoot_term(s)
	}
}

func (s *Root_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitRoot_term(s)
	}
}

func (p *ThreadComputationParser) Root_term() (localctx IRoot_termContext) {
	this := p
	_ = this

	localctx = NewRoot_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ThreadComputationParserRULE_root_term)

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

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserFUNC_NAME, ThreadComputationParserMOD:
		{
			p.SetState(20)
			p.Fun()
		}

	case ThreadComputationParserT__0:
		{
			p.SetState(21)
			p.Ns()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(22)
			p.Pos_term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INsContext is an interface to support dynamic dispatch.
type INsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNsname returns the nsname token.
	GetNsname() antlr.Token

	// SetNsname sets the nsname token.
	SetNsname(antlr.Token)

	// Get_ns_term returns the _ns_term rule contexts.
	Get_ns_term() INs_termContext

	// Set_ns_term sets the _ns_term rule contexts.
	Set_ns_term(INs_termContext)

	// GetParam returns the param rule context list.
	GetParam() []INs_termContext

	// SetParam sets the param rule context list.
	SetParam([]INs_termContext)

	// IsNsContext differentiates from other interfaces.
	IsNsContext()
}

type NsContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	nsname   antlr.Token
	_ns_term INs_termContext
	param    []INs_termContext
}

func NewEmptyNsContext() *NsContext {
	var p = new(NsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_ns
	return p
}

func (*NsContext) IsNsContext() {}

func NewNsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NsContext {
	var p = new(NsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_ns

	return p
}

func (s *NsContext) GetParser() antlr.Parser { return s.parser }

func (s *NsContext) GetNsname() antlr.Token { return s.nsname }

func (s *NsContext) SetNsname(v antlr.Token) { s.nsname = v }

func (s *NsContext) Get_ns_term() INs_termContext { return s._ns_term }

func (s *NsContext) Set_ns_term(v INs_termContext) { s._ns_term = v }

func (s *NsContext) GetParam() []INs_termContext { return s.param }

func (s *NsContext) SetParam(v []INs_termContext) { s.param = v }

func (s *NsContext) FUNC_NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserFUNC_NAME, 0)
}

func (s *NsContext) AllNs_term() []INs_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INs_termContext)(nil)).Elem())
	var tst = make([]INs_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INs_termContext)
		}
	}

	return tst
}

func (s *NsContext) Ns_term(i int) INs_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INs_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INs_termContext)
}

func (s *NsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterNs(s)
	}
}

func (s *NsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitNs(s)
	}
}

func (p *ThreadComputationParser) Ns() (localctx INsContext) {
	this := p
	_ = this

	localctx = NewNsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ThreadComputationParserRULE_ns)
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
		p.SetState(25)
		p.Match(ThreadComputationParserT__0)
	}
	{
		p.SetState(26)

		var _m = p.Match(ThreadComputationParserFUNC_NAME)

		localctx.(*NsContext).nsname = _m
	}
	{
		p.SetState(27)
		p.Match(ThreadComputationParserT__1)
	}
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__0)|(1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
		{
			p.SetState(28)

			var _x = p.Ns_term()

			localctx.(*NsContext)._ns_term = _x
		}
		localctx.(*NsContext).param = append(localctx.(*NsContext).param, localctx.(*NsContext)._ns_term)

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(34)
		p.Match(ThreadComputationParserT__2)
	}

	return localctx
}

// IFunContext is an interface to support dynamic dispatch.
type IFunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMod returns the mod token.
	GetMod() antlr.Token

	// GetFname returns the fname token.
	GetFname() antlr.Token

	// SetMod sets the mod token.
	SetMod(antlr.Token)

	// SetFname sets the fname token.
	SetFname(antlr.Token)

	// Get_fun_term returns the _fun_term rule contexts.
	Get_fun_term() IFun_termContext

	// Set_fun_term sets the _fun_term rule contexts.
	Set_fun_term(IFun_termContext)

	// GetParam returns the param rule context list.
	GetParam() []IFun_termContext

	// SetParam sets the param rule context list.
	SetParam([]IFun_termContext)

	// IsFunContext differentiates from other interfaces.
	IsFunContext()
}

type FunContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	mod       antlr.Token
	fname     antlr.Token
	_fun_term IFun_termContext
	param     []IFun_termContext
}

func NewEmptyFunContext() *FunContext {
	var p = new(FunContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_fun
	return p
}

func (*FunContext) IsFunContext() {}

func NewFunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunContext {
	var p = new(FunContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_fun

	return p
}

func (s *FunContext) GetParser() antlr.Parser { return s.parser }

func (s *FunContext) GetMod() antlr.Token { return s.mod }

func (s *FunContext) GetFname() antlr.Token { return s.fname }

func (s *FunContext) SetMod(v antlr.Token) { s.mod = v }

func (s *FunContext) SetFname(v antlr.Token) { s.fname = v }

func (s *FunContext) Get_fun_term() IFun_termContext { return s._fun_term }

func (s *FunContext) Set_fun_term(v IFun_termContext) { s._fun_term = v }

func (s *FunContext) GetParam() []IFun_termContext { return s.param }

func (s *FunContext) SetParam(v []IFun_termContext) { s.param = v }

func (s *FunContext) FUNC_NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserFUNC_NAME, 0)
}

func (s *FunContext) MOD() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserMOD, 0)
}

func (s *FunContext) AllFun_term() []IFun_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFun_termContext)(nil)).Elem())
	var tst = make([]IFun_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFun_termContext)
		}
	}

	return tst
}

func (s *FunContext) Fun_term(i int) IFun_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFun_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFun_termContext)
}

func (s *FunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterFun(s)
	}
}

func (s *FunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitFun(s)
	}
}

func (p *ThreadComputationParser) Fun() (localctx IFunContext) {
	this := p
	_ = this

	localctx = NewFunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ThreadComputationParserRULE_fun)
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
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThreadComputationParserMOD {
		{
			p.SetState(36)

			var _m = p.Match(ThreadComputationParserMOD)

			localctx.(*FunContext).mod = _m
		}

	}
	{
		p.SetState(39)

		var _m = p.Match(ThreadComputationParserFUNC_NAME)

		localctx.(*FunContext).fname = _m
	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(40)
			p.Match(ThreadComputationParserT__0)
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
			{
				p.SetState(41)

				var _x = p.Fun_term()

				localctx.(*FunContext)._fun_term = _x
			}
			localctx.(*FunContext).param = append(localctx.(*FunContext).param, localctx.(*FunContext)._fun_term)

			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(47)
			p.Match(ThreadComputationParserT__3)
		}

	}

	return localctx
}

// INs_termContext is an interface to support dynamic dispatch.
type INs_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNs_termContext differentiates from other interfaces.
	IsNs_termContext()
}

type Ns_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNs_termContext() *Ns_termContext {
	var p = new(Ns_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_ns_term
	return p
}

func (*Ns_termContext) IsNs_termContext() {}

func NewNs_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ns_termContext {
	var p = new(Ns_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_ns_term

	return p
}

func (s *Ns_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Ns_termContext) Fun() IFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *Ns_termContext) Ns() INsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *Ns_termContext) Pos_term() IPos_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPos_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPos_termContext)
}

func (s *Ns_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ns_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ns_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterNs_term(s)
	}
}

func (s *Ns_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitNs_term(s)
	}
}

func (p *ThreadComputationParser) Ns_term() (localctx INs_termContext) {
	this := p
	_ = this

	localctx = NewNs_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ThreadComputationParserRULE_ns_term)

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
	p.SetState(53)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserFUNC_NAME, ThreadComputationParserMOD:
		{
			p.SetState(50)
			p.Fun()
		}

	case ThreadComputationParserT__0:
		{
			p.SetState(51)
			p.Ns()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(52)
			p.Pos_term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFun_termContext is an interface to support dynamic dispatch.
type IFun_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFun_termContext differentiates from other interfaces.
	IsFun_termContext()
}

type Fun_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFun_termContext() *Fun_termContext {
	var p = new(Fun_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_fun_term
	return p
}

func (*Fun_termContext) IsFun_termContext() {}

func NewFun_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fun_termContext {
	var p = new(Fun_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_fun_term

	return p
}

func (s *Fun_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Fun_termContext) Fun() IFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *Fun_termContext) Pos_term() IPos_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPos_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPos_termContext)
}

func (s *Fun_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fun_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fun_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterFun_term(s)
	}
}

func (s *Fun_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitFun_term(s)
	}
}

func (p *ThreadComputationParser) Fun_term() (localctx IFun_termContext) {
	this := p
	_ = this

	localctx = NewFun_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ThreadComputationParserRULE_fun_term)

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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserFUNC_NAME, ThreadComputationParserMOD:
		{
			p.SetState(55)
			p.Fun()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(56)
			p.Pos_term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPos_termContext is an interface to support dynamic dispatch.
type IPos_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetPname returns the pname token.
	GetPname() antlr.Token

	// SetPname sets the pname token.
	SetPname(antlr.Token)

	// IsPos_termContext differentiates from other interfaces.
	IsPos_termContext()
}

type Pos_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	pname  antlr.Token
}

func NewEmptyPos_termContext() *Pos_termContext {
	var p = new(Pos_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_pos_term
	return p
}

func (*Pos_termContext) IsPos_termContext() {}

func NewPos_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pos_termContext {
	var p = new(Pos_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_pos_term

	return p
}

func (s *Pos_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Pos_termContext) GetPname() antlr.Token { return s.pname }

func (s *Pos_termContext) SetPname(v antlr.Token) { s.pname = v }

func (s *Pos_termContext) FUNC_NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserFUNC_NAME, 0)
}

func (s *Pos_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pos_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pos_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterPos_term(s)
	}
}

func (s *Pos_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitPos_term(s)
	}
}

func (p *ThreadComputationParser) Pos_term() (localctx IPos_termContext) {
	this := p
	_ = this

	localctx = NewPos_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ThreadComputationParserRULE_pos_term)

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
		p.SetState(59)
		p.Match(ThreadComputationParserT__4)
	}
	{
		p.SetState(60)

		var _m = p.Match(ThreadComputationParserFUNC_NAME)

		localctx.(*Pos_termContext).pname = _m
	}

	return localctx
}
