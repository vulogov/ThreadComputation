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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 24, 126,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 7,
	2, 36, 10, 2, 12, 2, 14, 2, 39, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 47, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 52, 10, 4, 12, 4, 14, 4, 55,
	11, 4, 3, 4, 5, 4, 58, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 64, 10, 5,
	12, 5, 14, 5, 67, 11, 5, 3, 5, 5, 5, 70, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 5, 7, 77, 10, 7, 3, 8, 3, 8, 5, 8, 81, 10, 8, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 88, 10, 9, 3, 10, 3, 10, 7, 10, 92, 10, 10, 12, 10, 14,
	10, 95, 11, 10, 3, 10, 3, 10, 3, 11, 3, 11, 7, 11, 101, 10, 11, 12, 11,
	14, 11, 104, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17,
	5, 17, 124, 10, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 2, 5, 4, 2, 12, 12, 14, 14, 3, 2, 20, 21, 5, 2,
	15, 15, 17, 18, 20, 21, 2, 130, 2, 37, 3, 2, 2, 2, 4, 46, 3, 2, 2, 2, 6,
	48, 3, 2, 2, 2, 8, 59, 3, 2, 2, 2, 10, 71, 3, 2, 2, 2, 12, 76, 3, 2, 2,
	2, 14, 80, 3, 2, 2, 2, 16, 87, 3, 2, 2, 2, 18, 89, 3, 2, 2, 2, 20, 98,
	3, 2, 2, 2, 22, 107, 3, 2, 2, 2, 24, 109, 3, 2, 2, 2, 26, 111, 3, 2, 2,
	2, 28, 113, 3, 2, 2, 2, 30, 115, 3, 2, 2, 2, 32, 123, 3, 2, 2, 2, 34, 36,
	5, 4, 3, 2, 35, 34, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2,
	37, 38, 3, 2, 2, 2, 38, 3, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 47, 5, 6,
	4, 2, 41, 47, 5, 8, 5, 2, 42, 47, 5, 32, 17, 2, 43, 47, 5, 10, 6, 2, 44,
	47, 5, 18, 10, 2, 45, 47, 5, 20, 11, 2, 46, 40, 3, 2, 2, 2, 46, 41, 3,
	2, 2, 2, 46, 42, 3, 2, 2, 2, 46, 43, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46,
	45, 3, 2, 2, 2, 47, 5, 3, 2, 2, 2, 48, 57, 9, 2, 2, 2, 49, 53, 7, 3, 2,
	2, 50, 52, 5, 14, 8, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2,
	56, 58, 7, 4, 2, 2, 57, 49, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 7, 3, 2,
	2, 2, 59, 60, 7, 5, 2, 2, 60, 69, 9, 2, 2, 2, 61, 65, 7, 3, 2, 2, 62, 64,
	5, 16, 9, 2, 63, 62, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2,
	65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 70, 7,
	4, 2, 2, 69, 61, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 9, 3, 2, 2, 2, 71,
	72, 7, 6, 2, 2, 72, 73, 7, 14, 2, 2, 73, 11, 3, 2, 2, 2, 74, 77, 5, 6,
	4, 2, 75, 77, 5, 32, 17, 2, 76, 74, 3, 2, 2, 2, 76, 75, 3, 2, 2, 2, 77,
	13, 3, 2, 2, 2, 78, 81, 5, 6, 4, 2, 79, 81, 5, 32, 17, 2, 80, 78, 3, 2,
	2, 2, 80, 79, 3, 2, 2, 2, 81, 15, 3, 2, 2, 2, 82, 88, 5, 6, 4, 2, 83, 88,
	5, 32, 17, 2, 84, 88, 5, 10, 6, 2, 85, 88, 5, 18, 10, 2, 86, 88, 5, 20,
	11, 2, 87, 82, 3, 2, 2, 2, 87, 83, 3, 2, 2, 2, 87, 84, 3, 2, 2, 2, 87,
	85, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 17, 3, 2, 2, 2, 89, 93, 7, 7, 2,
	2, 90, 92, 5, 12, 7, 2, 91, 90, 3, 2, 2, 2, 92, 95, 3, 2, 2, 2, 93, 91,
	3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 96, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2,
	96, 97, 7, 8, 2, 2, 97, 19, 3, 2, 2, 2, 98, 102, 7, 9, 2, 2, 99, 101, 5,
	30, 16, 2, 100, 99, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2,
	2, 102, 103, 3, 2, 2, 2, 103, 105, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105,
	106, 7, 10, 2, 2, 106, 21, 3, 2, 2, 2, 107, 108, 7, 15, 2, 2, 108, 23,
	3, 2, 2, 2, 109, 110, 7, 17, 2, 2, 110, 25, 3, 2, 2, 2, 111, 112, 7, 18,
	2, 2, 112, 27, 3, 2, 2, 2, 113, 114, 9, 3, 2, 2, 114, 29, 3, 2, 2, 2, 115,
	116, 7, 14, 2, 2, 116, 117, 7, 11, 2, 2, 117, 118, 9, 4, 2, 2, 118, 31,
	3, 2, 2, 2, 119, 124, 5, 22, 12, 2, 120, 124, 5, 24, 13, 2, 121, 124, 5,
	26, 14, 2, 122, 124, 5, 28, 15, 2, 123, 119, 3, 2, 2, 2, 123, 120, 3, 2,
	2, 2, 123, 121, 3, 2, 2, 2, 123, 122, 3, 2, 2, 2, 124, 33, 3, 2, 2, 2,
	14, 37, 46, 53, 57, 65, 69, 76, 80, 87, 93, 102, 123,
}
var literalNames = []string{
	"", "'['", "']'", "'@'", "'$'", "'('", "')'", "'{'", "'}'", "':'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "OPS", "OPS_START", "NAME", "INTEGER",
	"DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING", "OP", "TRUE", "FALSE", "BLOCK_COMMENT",
	"WhiteSpace", "NewLine",
}

var ruleNames = []string{
	"expressions", "root_term", "fun", "ufun", "vars", "dblock_term", "fun_term",
	"ufun_term", "dblock", "dmap", "integer_term", "float_term", "string_term",
	"boolean_term", "key_term", "term",
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
	ThreadComputationParserOPS             = 10
	ThreadComputationParserOPS_START       = 11
	ThreadComputationParserNAME            = 12
	ThreadComputationParserINTEGER         = 13
	ThreadComputationParserDECIMAL_INTEGER = 14
	ThreadComputationParserFLOAT_NUMBER    = 15
	ThreadComputationParserSTRING          = 16
	ThreadComputationParserOP              = 17
	ThreadComputationParserTRUE            = 18
	ThreadComputationParserFALSE           = 19
	ThreadComputationParserBLOCK_COMMENT   = 20
	ThreadComputationParserWhiteSpace      = 21
	ThreadComputationParserNewLine         = 22
)

// ThreadComputationParser rules.
const (
	ThreadComputationParserRULE_expressions  = 0
	ThreadComputationParserRULE_root_term    = 1
	ThreadComputationParserRULE_fun          = 2
	ThreadComputationParserRULE_ufun         = 3
	ThreadComputationParserRULE_vars         = 4
	ThreadComputationParserRULE_dblock_term  = 5
	ThreadComputationParserRULE_fun_term     = 6
	ThreadComputationParserRULE_ufun_term    = 7
	ThreadComputationParserRULE_dblock       = 8
	ThreadComputationParserRULE_dmap         = 9
	ThreadComputationParserRULE_integer_term = 10
	ThreadComputationParserRULE_float_term   = 11
	ThreadComputationParserRULE_string_term  = 12
	ThreadComputationParserRULE_boolean_term = 13
	ThreadComputationParserRULE_key_term     = 14
	ThreadComputationParserRULE_term         = 15
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__2)|(1<<ThreadComputationParserT__3)|(1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserT__6)|(1<<ThreadComputationParserOPS)|(1<<ThreadComputationParserNAME)|(1<<ThreadComputationParserINTEGER)|(1<<ThreadComputationParserFLOAT_NUMBER)|(1<<ThreadComputationParserSTRING)|(1<<ThreadComputationParserTRUE)|(1<<ThreadComputationParserFALSE))) != 0 {
		{
			p.SetState(32)
			p.Root_term()
		}

		p.SetState(37)
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

func (s *Root_termContext) Ufun() IUfunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUfunContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUfunContext)
}

func (s *Root_termContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Root_termContext) Vars() IVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
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
	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserOPS, ThreadComputationParserNAME:
		{
			p.SetState(38)
			p.Fun()
		}

	case ThreadComputationParserT__2:
		{
			p.SetState(39)
			p.Ufun()
		}

	case ThreadComputationParserINTEGER, ThreadComputationParserFLOAT_NUMBER, ThreadComputationParserSTRING, ThreadComputationParserTRUE, ThreadComputationParserFALSE:
		{
			p.SetState(40)
			p.Term()
		}

	case ThreadComputationParserT__3:
		{
			p.SetState(41)
			p.Vars()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(42)
			p.Dblock()
		}

	case ThreadComputationParserT__6:
		{
			p.SetState(43)
			p.Dmap()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunContext is an interface to support dynamic dispatch.
type IFunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFname returns the fname token.
	GetFname() antlr.Token

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

func (s *FunContext) GetFname() antlr.Token { return s.fname }

func (s *FunContext) SetFname(v antlr.Token) { s.fname = v }

func (s *FunContext) Get_fun_term() IFun_termContext { return s._fun_term }

func (s *FunContext) Set_fun_term(v IFun_termContext) { s._fun_term = v }

func (s *FunContext) GetParam() []IFun_termContext { return s.param }

func (s *FunContext) SetParam(v []IFun_termContext) { s.param = v }

func (s *FunContext) NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserNAME, 0)
}

func (s *FunContext) OPS() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserOPS, 0)
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
	p.EnterRule(localctx, 4, ThreadComputationParserRULE_fun)
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
		p.SetState(46)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*FunContext).fname = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ThreadComputationParserOPS || _la == ThreadComputationParserNAME) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*FunContext).fname = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThreadComputationParserT__0 {
		{
			p.SetState(47)
			p.Match(ThreadComputationParserT__0)
		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserOPS)|(1<<ThreadComputationParserNAME)|(1<<ThreadComputationParserINTEGER)|(1<<ThreadComputationParserFLOAT_NUMBER)|(1<<ThreadComputationParserSTRING)|(1<<ThreadComputationParserTRUE)|(1<<ThreadComputationParserFALSE))) != 0 {
			{
				p.SetState(48)

				var _x = p.Fun_term()

				localctx.(*FunContext)._fun_term = _x
			}
			localctx.(*FunContext).param = append(localctx.(*FunContext).param, localctx.(*FunContext)._fun_term)

			p.SetState(53)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(54)
			p.Match(ThreadComputationParserT__1)
		}

	}

	return localctx
}

// IUfunContext is an interface to support dynamic dispatch.
type IUfunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetFname returns the fname token.
	GetFname() antlr.Token

	// SetFname sets the fname token.
	SetFname(antlr.Token)

	// Get_ufun_term returns the _ufun_term rule contexts.
	Get_ufun_term() IUfun_termContext

	// Set_ufun_term sets the _ufun_term rule contexts.
	Set_ufun_term(IUfun_termContext)

	// GetParam returns the param rule context list.
	GetParam() []IUfun_termContext

	// SetParam sets the param rule context list.
	SetParam([]IUfun_termContext)

	// IsUfunContext differentiates from other interfaces.
	IsUfunContext()
}

type UfunContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	fname      antlr.Token
	_ufun_term IUfun_termContext
	param      []IUfun_termContext
}

func NewEmptyUfunContext() *UfunContext {
	var p = new(UfunContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_ufun
	return p
}

func (*UfunContext) IsUfunContext() {}

func NewUfunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UfunContext {
	var p = new(UfunContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_ufun

	return p
}

func (s *UfunContext) GetParser() antlr.Parser { return s.parser }

func (s *UfunContext) GetFname() antlr.Token { return s.fname }

func (s *UfunContext) SetFname(v antlr.Token) { s.fname = v }

func (s *UfunContext) Get_ufun_term() IUfun_termContext { return s._ufun_term }

func (s *UfunContext) Set_ufun_term(v IUfun_termContext) { s._ufun_term = v }

func (s *UfunContext) GetParam() []IUfun_termContext { return s.param }

func (s *UfunContext) SetParam(v []IUfun_termContext) { s.param = v }

func (s *UfunContext) NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserNAME, 0)
}

func (s *UfunContext) OPS() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserOPS, 0)
}

func (s *UfunContext) AllUfun_term() []IUfun_termContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUfun_termContext)(nil)).Elem())
	var tst = make([]IUfun_termContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUfun_termContext)
		}
	}

	return tst
}

func (s *UfunContext) Ufun_term(i int) IUfun_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUfun_termContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUfun_termContext)
}

func (s *UfunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UfunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UfunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterUfun(s)
	}
}

func (s *UfunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitUfun(s)
	}
}

func (p *ThreadComputationParser) Ufun() (localctx IUfunContext) {
	this := p
	_ = this

	localctx = NewUfunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ThreadComputationParserRULE_ufun)
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
		p.SetState(57)
		p.Match(ThreadComputationParserT__2)
	}
	{
		p.SetState(58)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*UfunContext).fname = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ThreadComputationParserOPS || _la == ThreadComputationParserNAME) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*UfunContext).fname = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThreadComputationParserT__0 {
		{
			p.SetState(59)
			p.Match(ThreadComputationParserT__0)
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__3)|(1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserT__6)|(1<<ThreadComputationParserOPS)|(1<<ThreadComputationParserNAME)|(1<<ThreadComputationParserINTEGER)|(1<<ThreadComputationParserFLOAT_NUMBER)|(1<<ThreadComputationParserSTRING)|(1<<ThreadComputationParserTRUE)|(1<<ThreadComputationParserFALSE))) != 0 {
			{
				p.SetState(60)

				var _x = p.Ufun_term()

				localctx.(*UfunContext)._ufun_term = _x
			}
			localctx.(*UfunContext).param = append(localctx.(*UfunContext).param, localctx.(*UfunContext)._ufun_term)

			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(66)
			p.Match(ThreadComputationParserT__1)
		}

	}

	return localctx
}

// IVarsContext is an interface to support dynamic dispatch.
type IVarsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVname returns the vname token.
	GetVname() antlr.Token

	// SetVname sets the vname token.
	SetVname(antlr.Token)

	// IsVarsContext differentiates from other interfaces.
	IsVarsContext()
}

type VarsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	vname  antlr.Token
}

func NewEmptyVarsContext() *VarsContext {
	var p = new(VarsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_vars
	return p
}

func (*VarsContext) IsVarsContext() {}

func NewVarsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarsContext {
	var p = new(VarsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_vars

	return p
}

func (s *VarsContext) GetParser() antlr.Parser { return s.parser }

func (s *VarsContext) GetVname() antlr.Token { return s.vname }

func (s *VarsContext) SetVname(v antlr.Token) { s.vname = v }

func (s *VarsContext) NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserNAME, 0)
}

func (s *VarsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterVars(s)
	}
}

func (s *VarsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitVars(s)
	}
}

func (p *ThreadComputationParser) Vars() (localctx IVarsContext) {
	this := p
	_ = this

	localctx = NewVarsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ThreadComputationParserRULE_vars)

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
		p.SetState(69)
		p.Match(ThreadComputationParserT__3)
	}
	{
		p.SetState(70)

		var _m = p.Match(ThreadComputationParserNAME)

		localctx.(*VarsContext).vname = _m
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

func (s *Dblock_termContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
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
	p.SetState(74)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserOPS, ThreadComputationParserNAME:
		{
			p.SetState(72)
			p.Fun()
		}

	case ThreadComputationParserINTEGER, ThreadComputationParserFLOAT_NUMBER, ThreadComputationParserSTRING, ThreadComputationParserTRUE, ThreadComputationParserFALSE:
		{
			p.SetState(73)
			p.Term()
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

func (s *Fun_termContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
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
	p.SetState(78)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserOPS, ThreadComputationParserNAME:
		{
			p.SetState(76)
			p.Fun()
		}

	case ThreadComputationParserINTEGER, ThreadComputationParserFLOAT_NUMBER, ThreadComputationParserSTRING, ThreadComputationParserTRUE, ThreadComputationParserFALSE:
		{
			p.SetState(77)
			p.Term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUfun_termContext is an interface to support dynamic dispatch.
type IUfun_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUfun_termContext differentiates from other interfaces.
	IsUfun_termContext()
}

type Ufun_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUfun_termContext() *Ufun_termContext {
	var p = new(Ufun_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_ufun_term
	return p
}

func (*Ufun_termContext) IsUfun_termContext() {}

func NewUfun_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ufun_termContext {
	var p = new(Ufun_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_ufun_term

	return p
}

func (s *Ufun_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Ufun_termContext) Fun() IFunContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *Ufun_termContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Ufun_termContext) Vars() IVarsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarsContext)
}

func (s *Ufun_termContext) Dblock() IDblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDblockContext)
}

func (s *Ufun_termContext) Dmap() IDmapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDmapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDmapContext)
}

func (s *Ufun_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ufun_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ufun_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterUfun_term(s)
	}
}

func (s *Ufun_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitUfun_term(s)
	}
}

func (p *ThreadComputationParser) Ufun_term() (localctx IUfun_termContext) {
	this := p
	_ = this

	localctx = NewUfun_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ThreadComputationParserRULE_ufun_term)

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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserOPS, ThreadComputationParserNAME:
		{
			p.SetState(80)
			p.Fun()
		}

	case ThreadComputationParserINTEGER, ThreadComputationParserFLOAT_NUMBER, ThreadComputationParserSTRING, ThreadComputationParserTRUE, ThreadComputationParserFALSE:
		{
			p.SetState(81)
			p.Term()
		}

	case ThreadComputationParserT__3:
		{
			p.SetState(82)
			p.Vars()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(83)
			p.Dblock()
		}

	case ThreadComputationParserT__6:
		{
			p.SetState(84)
			p.Dmap()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDblockContext is an interface to support dynamic dispatch.
type IDblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (s *DblockContext) Get_dblock_term() IDblock_termContext { return s._dblock_term }

func (s *DblockContext) Set_dblock_term(v IDblock_termContext) { s._dblock_term = v }

func (s *DblockContext) GetParam() []IDblock_termContext { return s.param }

func (s *DblockContext) SetParam(v []IDblock_termContext) { s.param = v }

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
	{
		p.SetState(87)
		p.Match(ThreadComputationParserT__4)
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserOPS)|(1<<ThreadComputationParserNAME)|(1<<ThreadComputationParserINTEGER)|(1<<ThreadComputationParserFLOAT_NUMBER)|(1<<ThreadComputationParserSTRING)|(1<<ThreadComputationParserTRUE)|(1<<ThreadComputationParserFALSE))) != 0 {
		{
			p.SetState(88)

			var _x = p.Dblock_term()

			localctx.(*DblockContext)._dblock_term = _x
		}
		localctx.(*DblockContext).param = append(localctx.(*DblockContext).param, localctx.(*DblockContext)._dblock_term)

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(94)
		p.Match(ThreadComputationParserT__5)
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
		p.SetState(96)
		p.Match(ThreadComputationParserT__6)
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThreadComputationParserNAME {
		{
			p.SetState(97)

			var _x = p.Key_term()

			localctx.(*DmapContext)._key_term = _x
		}
		localctx.(*DmapContext).param = append(localctx.(*DmapContext).param, localctx.(*DmapContext)._key_term)

		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(ThreadComputationParserT__7)
	}

	return localctx
}

// IInteger_termContext is an interface to support dynamic dispatch.
type IInteger_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// IsInteger_termContext differentiates from other interfaces.
	IsInteger_termContext()
}

type Integer_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	VALUE  antlr.Token
}

func NewEmptyInteger_termContext() *Integer_termContext {
	var p = new(Integer_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_integer_term
	return p
}

func (*Integer_termContext) IsInteger_termContext() {}

func NewInteger_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_termContext {
	var p = new(Integer_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_integer_term

	return p
}

func (s *Integer_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Integer_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Integer_termContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserINTEGER, 0)
}

func (s *Integer_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterInteger_term(s)
	}
}

func (s *Integer_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitInteger_term(s)
	}
}

func (p *ThreadComputationParser) Integer_term() (localctx IInteger_termContext) {
	this := p
	_ = this

	localctx = NewInteger_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ThreadComputationParserRULE_integer_term)

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
		p.SetState(105)

		var _m = p.Match(ThreadComputationParserINTEGER)

		localctx.(*Integer_termContext).VALUE = _m
	}

	return localctx
}

// IFloat_termContext is an interface to support dynamic dispatch.
type IFloat_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// IsFloat_termContext differentiates from other interfaces.
	IsFloat_termContext()
}

type Float_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	VALUE  antlr.Token
}

func NewEmptyFloat_termContext() *Float_termContext {
	var p = new(Float_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_float_term
	return p
}

func (*Float_termContext) IsFloat_termContext() {}

func NewFloat_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_termContext {
	var p = new(Float_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_float_term

	return p
}

func (s *Float_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Float_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Float_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Float_termContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserFLOAT_NUMBER, 0)
}

func (s *Float_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterFloat_term(s)
	}
}

func (s *Float_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitFloat_term(s)
	}
}

func (p *ThreadComputationParser) Float_term() (localctx IFloat_termContext) {
	this := p
	_ = this

	localctx = NewFloat_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ThreadComputationParserRULE_float_term)

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
		p.SetState(107)

		var _m = p.Match(ThreadComputationParserFLOAT_NUMBER)

		localctx.(*Float_termContext).VALUE = _m
	}

	return localctx
}

// IString_termContext is an interface to support dynamic dispatch.
type IString_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// IsString_termContext differentiates from other interfaces.
	IsString_termContext()
}

type String_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	VALUE  antlr.Token
}

func NewEmptyString_termContext() *String_termContext {
	var p = new(String_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_string_term
	return p
}

func (*String_termContext) IsString_termContext() {}

func NewString_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_termContext {
	var p = new(String_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_string_term

	return p
}

func (s *String_termContext) GetParser() antlr.Parser { return s.parser }

func (s *String_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *String_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *String_termContext) STRING() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserSTRING, 0)
}

func (s *String_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterString_term(s)
	}
}

func (s *String_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitString_term(s)
	}
}

func (p *ThreadComputationParser) String_term() (localctx IString_termContext) {
	this := p
	_ = this

	localctx = NewString_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ThreadComputationParserRULE_string_term)

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
		p.SetState(109)

		var _m = p.Match(ThreadComputationParserSTRING)

		localctx.(*String_termContext).VALUE = _m
	}

	return localctx
}

// IBoolean_termContext is an interface to support dynamic dispatch.
type IBoolean_termContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// IsBoolean_termContext differentiates from other interfaces.
	IsBoolean_termContext()
}

type Boolean_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	VALUE  antlr.Token
}

func NewEmptyBoolean_termContext() *Boolean_termContext {
	var p = new(Boolean_termContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_boolean_term
	return p
}

func (*Boolean_termContext) IsBoolean_termContext() {}

func NewBoolean_termContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_termContext {
	var p = new(Boolean_termContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_boolean_term

	return p
}

func (s *Boolean_termContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Boolean_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Boolean_termContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserTRUE, 0)
}

func (s *Boolean_termContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserFALSE, 0)
}

func (s *Boolean_termContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_termContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_termContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterBoolean_term(s)
	}
}

func (s *Boolean_termContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitBoolean_term(s)
	}
}

func (p *ThreadComputationParser) Boolean_term() (localctx IBoolean_termContext) {
	this := p
	_ = this

	localctx = NewBoolean_termContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ThreadComputationParserRULE_boolean_term)
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
		p.SetState(111)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Boolean_termContext).VALUE = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == ThreadComputationParserTRUE || _la == ThreadComputationParserFALSE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Boolean_termContext).VALUE = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

	// GetVALUE returns the VALUE token.
	GetVALUE() antlr.Token

	// SetKEY sets the KEY token.
	SetKEY(antlr.Token)

	// SetVALUE sets the VALUE token.
	SetVALUE(antlr.Token)

	// IsKey_termContext differentiates from other interfaces.
	IsKey_termContext()
}

type Key_termContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	KEY    antlr.Token
	VALUE  antlr.Token
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

func (s *Key_termContext) GetVALUE() antlr.Token { return s.VALUE }

func (s *Key_termContext) SetKEY(v antlr.Token) { s.KEY = v }

func (s *Key_termContext) SetVALUE(v antlr.Token) { s.VALUE = v }

func (s *Key_termContext) NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserNAME, 0)
}

func (s *Key_termContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserINTEGER, 0)
}

func (s *Key_termContext) FLOAT_NUMBER() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserFLOAT_NUMBER, 0)
}

func (s *Key_termContext) STRING() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserSTRING, 0)
}

func (s *Key_termContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserTRUE, 0)
}

func (s *Key_termContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserFALSE, 0)
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
	p.EnterRule(localctx, 28, ThreadComputationParserRULE_key_term)
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
		p.SetState(113)

		var _m = p.Match(ThreadComputationParserNAME)

		localctx.(*Key_termContext).KEY = _m
	}
	{
		p.SetState(114)
		p.Match(ThreadComputationParserT__8)
	}
	{
		p.SetState(115)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Key_termContext).VALUE = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserINTEGER)|(1<<ThreadComputationParserFLOAT_NUMBER)|(1<<ThreadComputationParserSTRING)|(1<<ThreadComputationParserTRUE)|(1<<ThreadComputationParserFALSE))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Key_termContext).VALUE = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Integer_term() IInteger_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_termContext)
}

func (s *TermContext) Float_term() IFloat_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloat_termContext)
}

func (s *TermContext) String_term() IString_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_termContext)
}

func (s *TermContext) Boolean_term() IBoolean_termContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_termContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_termContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *ThreadComputationParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ThreadComputationParserRULE_term)

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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserINTEGER:
		{
			p.SetState(117)
			p.Integer_term()
		}

	case ThreadComputationParserFLOAT_NUMBER:
		{
			p.SetState(118)
			p.Float_term()
		}

	case ThreadComputationParserSTRING:
		{
			p.SetState(119)
			p.String_term()
		}

	case ThreadComputationParserTRUE, ThreadComputationParserFALSE:
		{
			p.SetState(120)
			p.Boolean_term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
