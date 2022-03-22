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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 108,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	7, 2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 36, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 42, 10, 4, 12, 4, 14, 4,
	45, 11, 4, 3, 4, 3, 4, 3, 5, 5, 5, 50, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5, 55,
	10, 5, 12, 5, 14, 5, 58, 11, 5, 3, 5, 5, 5, 61, 10, 5, 3, 6, 3, 6, 3, 6,
	5, 6, 66, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 72, 10, 7, 3, 8, 3, 8, 3,
	8, 5, 8, 77, 10, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 5, 10, 84, 10, 10,
	3, 10, 3, 10, 7, 10, 88, 10, 10, 12, 10, 14, 10, 91, 11, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 7, 11, 97, 10, 11, 12, 11, 14, 11, 100, 11, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12,
	14, 16, 18, 20, 22, 2, 2, 2, 115, 2, 27, 3, 2, 2, 2, 4, 35, 3, 2, 2, 2,
	6, 37, 3, 2, 2, 2, 8, 49, 3, 2, 2, 2, 10, 65, 3, 2, 2, 2, 12, 71, 3, 2,
	2, 2, 14, 76, 3, 2, 2, 2, 16, 78, 3, 2, 2, 2, 18, 83, 3, 2, 2, 2, 20, 94,
	3, 2, 2, 2, 22, 103, 3, 2, 2, 2, 24, 26, 5, 4, 3, 2, 25, 24, 3, 2, 2, 2,
	26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 3, 3, 2,
	2, 2, 29, 27, 3, 2, 2, 2, 30, 36, 5, 8, 5, 2, 31, 36, 5, 18, 10, 2, 32,
	36, 5, 20, 11, 2, 33, 36, 5, 6, 4, 2, 34, 36, 5, 16, 9, 2, 35, 30, 3, 2,
	2, 2, 35, 31, 3, 2, 2, 2, 35, 32, 3, 2, 2, 2, 35, 33, 3, 2, 2, 2, 35, 34,
	3, 2, 2, 2, 36, 5, 3, 2, 2, 2, 37, 38, 7, 3, 2, 2, 38, 39, 7, 12, 2, 2,
	39, 43, 7, 4, 2, 2, 40, 42, 5, 10, 6, 2, 41, 40, 3, 2, 2, 2, 42, 45, 3,
	2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 46, 3, 2, 2, 2, 45,
	43, 3, 2, 2, 2, 46, 47, 7, 5, 2, 2, 47, 7, 3, 2, 2, 2, 48, 50, 7, 20, 2,
	2, 49, 48, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 51, 3, 2, 2, 2, 51, 60,
	7, 12, 2, 2, 52, 56, 7, 3, 2, 2, 53, 55, 5, 14, 8, 2, 54, 53, 3, 2, 2,
	2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 59,
	3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 61, 7, 6, 2, 2, 60, 52, 3, 2, 2, 2,
	60, 61, 3, 2, 2, 2, 61, 9, 3, 2, 2, 2, 62, 66, 5, 8, 5, 2, 63, 66, 5, 6,
	4, 2, 64, 66, 5, 20, 11, 2, 65, 62, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65,
	64, 3, 2, 2, 2, 66, 11, 3, 2, 2, 2, 67, 72, 5, 8, 5, 2, 68, 72, 5, 18,
	10, 2, 69, 72, 5, 20, 11, 2, 70, 72, 5, 16, 9, 2, 71, 67, 3, 2, 2, 2, 71,
	68, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 70, 3, 2, 2, 2, 72, 13, 3, 2, 2,
	2, 73, 77, 5, 8, 5, 2, 74, 77, 5, 20, 11, 2, 75, 77, 5, 16, 9, 2, 76, 73,
	3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 75, 3, 2, 2, 2, 77, 15, 3, 2, 2, 2,
	78, 79, 7, 7, 2, 2, 79, 80, 7, 12, 2, 2, 80, 17, 3, 2, 2, 2, 81, 82, 7,
	4, 2, 2, 82, 84, 7, 14, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84,
	85, 3, 2, 2, 2, 85, 89, 7, 8, 2, 2, 86, 88, 5, 12, 7, 2, 87, 86, 3, 2,
	2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 92,
	3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 7, 9, 2, 2, 93, 19, 3, 2, 2, 2,
	94, 98, 7, 10, 2, 2, 95, 97, 5, 22, 12, 2, 96, 95, 3, 2, 2, 2, 97, 100,
	3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 101, 3, 2, 2, 2,
	100, 98, 3, 2, 2, 2, 101, 102, 7, 11, 2, 2, 102, 21, 3, 2, 2, 2, 103, 104,
	7, 14, 2, 2, 104, 105, 7, 4, 2, 2, 105, 106, 5, 14, 8, 2, 106, 23, 3, 2,
	2, 2, 14, 27, 35, 43, 49, 56, 60, 65, 71, 76, 83, 89, 98,
}
var literalNames = []string{
	"", "'['", "':'", "';;'", "']'", "'#'", "'('", "')'", "'{'", "'}'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "FUNC_NAME", "OPS", "NAME", "INTEGER",
	"DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING", "OP", "MOD", "BLOCK_COMMENT",
	"WhiteSpace", "NewLine",
}

var ruleNames = []string{
	"expressions", "root_term", "ns", "fun", "ns_term", "dblock_term", "fun_term",
	"pos_term", "dblock", "dmap", "key_term",
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
	ThreadComputationParserT__5            = 6
	ThreadComputationParserT__6            = 7
	ThreadComputationParserT__7            = 8
	ThreadComputationParserT__8            = 9
	ThreadComputationParserFUNC_NAME       = 10
	ThreadComputationParserOPS             = 11
	ThreadComputationParserNAME            = 12
	ThreadComputationParserINTEGER         = 13
	ThreadComputationParserDECIMAL_INTEGER = 14
	ThreadComputationParserFLOAT_NUMBER    = 15
	ThreadComputationParserSTRING          = 16
	ThreadComputationParserOP              = 17
	ThreadComputationParserMOD             = 18
	ThreadComputationParserBLOCK_COMMENT   = 19
	ThreadComputationParserWhiteSpace      = 20
	ThreadComputationParserNewLine         = 21
)

// ThreadComputationParser rules.
const (
	ThreadComputationParserRULE_expressions = 0
	ThreadComputationParserRULE_root_term   = 1
	ThreadComputationParserRULE_ns          = 2
	ThreadComputationParserRULE_fun         = 3
	ThreadComputationParserRULE_ns_term     = 4
	ThreadComputationParserRULE_dblock_term = 5
	ThreadComputationParserRULE_fun_term    = 6
	ThreadComputationParserRULE_pos_term    = 7
	ThreadComputationParserRULE_dblock      = 8
	ThreadComputationParserRULE_dmap        = 9
	ThreadComputationParserRULE_key_term    = 10
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
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__0)|(1<<ThreadComputationParserT__1)|(1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserT__5)|(1<<ThreadComputationParserT__7)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
		{
			p.SetState(22)
			p.Root_term()
		}

		p.SetState(27)
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

func (s *Root_termContext) Dblock() IDblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDblockContext)
}

func (s *Root_termContext) Dmap() IDmapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDmapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDmapContext)
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
	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserFUNC_NAME, ThreadComputationParserMOD:
		{
			p.SetState(28)
			p.Fun()
		}

	case ThreadComputationParserT__1, ThreadComputationParserT__5:
		{
			p.SetState(29)
			p.Dblock()
		}

	case ThreadComputationParserT__7:
		{
			p.SetState(30)
			p.Dmap()
		}

	case ThreadComputationParserT__0:
		{
			p.SetState(31)
			p.Ns()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(32)
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

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

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
	name     antlr.Token
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

func (s *NsContext) GetName() antlr.Token { return s.name }

func (s *NsContext) SetName(v antlr.Token) { s.name = v }

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
		p.SetState(35)
		p.Match(ThreadComputationParserT__0)
	}
	{
		p.SetState(36)

		var _m = p.Match(ThreadComputationParserFUNC_NAME)

		localctx.(*NsContext).name = _m
	}
	{
		p.SetState(37)
		p.Match(ThreadComputationParserT__1)
	}
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__0)|(1<<ThreadComputationParserT__7)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
		{
			p.SetState(38)

			var _x = p.Ns_term()

			localctx.(*NsContext)._ns_term = _x
		}
		localctx.(*NsContext).param = append(localctx.(*NsContext).param, localctx.(*NsContext)._ns_term)

		p.SetState(43)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(44)
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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThreadComputationParserMOD {
		{
			p.SetState(46)

			var _m = p.Match(ThreadComputationParserMOD)

			localctx.(*FunContext).mod = _m
		}

	}
	{
		p.SetState(49)

		var _m = p.Match(ThreadComputationParserFUNC_NAME)

		localctx.(*FunContext).fname = _m
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(50)
			p.Match(ThreadComputationParserT__0)
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserT__7)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
			{
				p.SetState(51)

				var _x = p.Fun_term()

				localctx.(*FunContext)._fun_term = _x
			}
			localctx.(*FunContext).param = append(localctx.(*FunContext).param, localctx.(*FunContext)._fun_term)

			p.SetState(56)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(57)
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

func (s *Ns_termContext) Dmap() IDmapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDmapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDmapContext)
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
	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserFUNC_NAME, ThreadComputationParserMOD:
		{
			p.SetState(60)
			p.Fun()
		}

	case ThreadComputationParserT__0:
		{
			p.SetState(61)
			p.Ns()
		}

	case ThreadComputationParserT__7:
		{
			p.SetState(62)
			p.Dmap()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDblock_termContext is an interface to support dynamic dispatch.
type IDblock_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDblock_termContext differentiates from other interfaces.
	IsDblock_termContext()
}

type Dblock_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDblock_termContext() *Dblock_termContext {
	var p = new(Dblock_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_dblock_term
	return p
}

func (*Dblock_termContext) IsDblock_termContext() {}

func NewDblock_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dblock_termContext {
	var p = new(Dblock_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_dblock_term

	return p
}

func (s *Dblock_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Dblock_termContext) Fun() IFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *Dblock_termContext) Dblock() IDblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDblockContext)
}

func (s *Dblock_termContext) Dmap() IDmapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDmapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDmapContext)
}

func (s *Dblock_termContext) Pos_term() IPos_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPos_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPos_termContext)
}

func (s *Dblock_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dblock_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dblock_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterDblock_term(s)
	}
}

func (s *Dblock_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitDblock_term(s)
	}
}

func (p *ThreadComputationParser) Dblock_term() (localctx IDblock_termContext) {
	this := p
	_ = this

	localctx = NewDblock_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ThreadComputationParserRULE_dblock_term)

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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserFUNC_NAME, ThreadComputationParserMOD:
		{
			p.SetState(65)
			p.Fun()
		}

	case ThreadComputationParserT__1, ThreadComputationParserT__5:
		{
			p.SetState(66)
			p.Dblock()
		}

	case ThreadComputationParserT__7:
		{
			p.SetState(67)
			p.Dmap()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(68)
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

func (s *Fun_termContext) Dmap() IDmapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDmapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDmapContext)
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
	p.EnterRule(localctx, 12, ThreadComputationParserRULE_fun_term)

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
	p.SetState(74)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserFUNC_NAME, ThreadComputationParserMOD:
		{
			p.SetState(71)
			p.Fun()
		}

	case ThreadComputationParserT__7:
		{
			p.SetState(72)
			p.Dmap()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(73)
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
	p.EnterRule(localctx, 14, ThreadComputationParserRULE_pos_term)

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
		p.SetState(76)
		p.Match(ThreadComputationParserT__4)
	}
	{
		p.SetState(77)

		var _m = p.Match(ThreadComputationParserFUNC_NAME)

		localctx.(*Pos_termContext).pname = _m
	}

	return localctx
}

// IDblockContext is an interface to support dynamic dispatch.
type IDblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBname returns the bname token.
	GetBname() antlr.Token

	// SetBname sets the bname token.
	SetBname(antlr.Token)

	// Get_dblock_term returns the _dblock_term rule contexts.
	Get_dblock_term() IDblock_termContext

	// Set_dblock_term sets the _dblock_term rule contexts.
	Set_dblock_term(IDblock_termContext)

	// GetParam returns the param rule context list.
	GetParam() []IDblock_termContext

	// SetParam sets the param rule context list.
	SetParam([]IDblock_termContext)

	// IsDblockContext differentiates from other interfaces.
	IsDblockContext()
}

type DblockContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	bname        antlr.Token
	_dblock_term IDblock_termContext
	param        []IDblock_termContext
}

func NewEmptyDblockContext() *DblockContext {
	var p = new(DblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_dblock
	return p
}

func (*DblockContext) IsDblockContext() {}

func NewDblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DblockContext {
	var p = new(DblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_dblock

	return p
}

func (s *DblockContext) GetParser() antlr.Parser { return s.parser }

func (s *DblockContext) GetBname() antlr.Token { return s.bname }

func (s *DblockContext) SetBname(v antlr.Token) { s.bname = v }

func (s *DblockContext) Get_dblock_term() IDblock_termContext { return s._dblock_term }

func (s *DblockContext) Set_dblock_term(v IDblock_termContext) { s._dblock_term = v }

func (s *DblockContext) GetParam() []IDblock_termContext { return s.param }

func (s *DblockContext) SetParam(v []IDblock_termContext) { s.param = v }

func (s *DblockContext) NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserNAME, 0)
}

func (s *DblockContext) AllDblock_term() []IDblock_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDblock_termContext)(nil)).Elem())
	var tst = make([]IDblock_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDblock_termContext)
		}
	}

	return tst
}

func (s *DblockContext) Dblock_term(i int) IDblock_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDblock_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDblock_termContext)
}

func (s *DblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterDblock(s)
	}
}

func (s *DblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitDblock(s)
	}
}

func (p *ThreadComputationParser) Dblock() (localctx IDblockContext) {
	this := p
	_ = this

	localctx = NewDblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ThreadComputationParserRULE_dblock)
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThreadComputationParserT__1 {
		{
			p.SetState(79)
			p.Match(ThreadComputationParserT__1)
		}
		{
			p.SetState(80)

			var _m = p.Match(ThreadComputationParserNAME)

			localctx.(*DblockContext).bname = _m
		}

	}
	{
		p.SetState(83)
		p.Match(ThreadComputationParserT__5)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__1)|(1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserT__5)|(1<<ThreadComputationParserT__7)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
		{
			p.SetState(84)

			var _x = p.Dblock_term()

			localctx.(*DblockContext)._dblock_term = _x
		}
		localctx.(*DblockContext).param = append(localctx.(*DblockContext).param, localctx.(*DblockContext)._dblock_term)

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(ThreadComputationParserT__6)
	}

	return localctx
}

// IDmapContext is an interface to support dynamic dispatch.
type IDmapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_key_term returns the _key_term rule contexts.
	Get_key_term() IKey_termContext

	// Set_key_term sets the _key_term rule contexts.
	Set_key_term(IKey_termContext)

	// GetParam returns the param rule context list.
	GetParam() []IKey_termContext

	// SetParam sets the param rule context list.
	SetParam([]IKey_termContext)

	// IsDmapContext differentiates from other interfaces.
	IsDmapContext()
}

type DmapContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	_key_term IKey_termContext
	param     []IKey_termContext
}

func NewEmptyDmapContext() *DmapContext {
	var p = new(DmapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_dmap
	return p
}

func (*DmapContext) IsDmapContext() {}

func NewDmapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DmapContext {
	var p = new(DmapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_dmap

	return p
}

func (s *DmapContext) GetParser() antlr.Parser { return s.parser }

func (s *DmapContext) Get_key_term() IKey_termContext { return s._key_term }

func (s *DmapContext) Set_key_term(v IKey_termContext) { s._key_term = v }

func (s *DmapContext) GetParam() []IKey_termContext { return s.param }

func (s *DmapContext) SetParam(v []IKey_termContext) { s.param = v }

func (s *DmapContext) AllKey_term() []IKey_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKey_termContext)(nil)).Elem())
	var tst = make([]IKey_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKey_termContext)
		}
	}

	return tst
}

func (s *DmapContext) Key_term(i int) IKey_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKey_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKey_termContext)
}

func (s *DmapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DmapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DmapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterDmap(s)
	}
}

func (s *DmapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitDmap(s)
	}
}

func (p *ThreadComputationParser) Dmap() (localctx IDmapContext) {
	this := p
	_ = this

	localctx = NewDmapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ThreadComputationParserRULE_dmap)
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
		p.SetState(92)
		p.Match(ThreadComputationParserT__7)
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThreadComputationParserNAME {
		{
			p.SetState(93)

			var _x = p.Key_term()

			localctx.(*DmapContext)._key_term = _x
		}
		localctx.(*DmapContext).param = append(localctx.(*DmapContext).param, localctx.(*DmapContext)._key_term)

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(ThreadComputationParserT__8)
	}

	return localctx
}

// IKey_termContext is an interface to support dynamic dispatch.
type IKey_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKEY returns the KEY token.
	GetKEY() antlr.Token

	// SetKEY sets the KEY token.
	SetKEY(antlr.Token)

	// GetVALUE returns the VALUE rule contexts.
	GetVALUE() IFun_termContext

	// SetVALUE sets the VALUE rule contexts.
	SetVALUE(IFun_termContext)

	// IsKey_termContext differentiates from other interfaces.
	IsKey_termContext()
}

type Key_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	KEY    antlr.Token
	VALUE  IFun_termContext
}

func NewEmptyKey_termContext() *Key_termContext {
	var p = new(Key_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_key_term
	return p
}

func (*Key_termContext) IsKey_termContext() {}

func NewKey_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Key_termContext {
	var p = new(Key_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_key_term

	return p
}

func (s *Key_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Key_termContext) GetKEY() antlr.Token { return s.KEY }

func (s *Key_termContext) SetKEY(v antlr.Token) { s.KEY = v }

func (s *Key_termContext) GetVALUE() IFun_termContext { return s.VALUE }

func (s *Key_termContext) SetVALUE(v IFun_termContext) { s.VALUE = v }

func (s *Key_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserNAME, 0)
}

func (s *Key_termContext) Fun_term() IFun_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFun_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFun_termContext)
}

func (s *Key_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Key_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Key_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterKey_term(s)
	}
}

func (s *Key_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitKey_term(s)
	}
}

func (p *ThreadComputationParser) Key_term() (localctx IKey_termContext) {
	this := p
	_ = this

	localctx = NewKey_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ThreadComputationParserRULE_key_term)

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
		p.SetState(101)

		var _m = p.Match(ThreadComputationParserNAME)

		localctx.(*Key_termContext).KEY = _m
	}
	{
		p.SetState(102)
		p.Match(ThreadComputationParserT__1)
	}
	{
		p.SetState(103)

		var _x = p.Fun_term()

		localctx.(*Key_termContext).VALUE = _x
	}

	return localctx
}
