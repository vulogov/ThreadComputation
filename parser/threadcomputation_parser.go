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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 78, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2, 7,
	2, 26, 10, 2, 12, 2, 14, 2, 29, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 35,
	10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 40, 10, 4, 12, 4, 14, 4, 43, 11, 4, 3, 4,
	5, 4, 46, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 53, 10, 6, 3, 7, 3,
	7, 7, 7, 57, 10, 7, 12, 7, 14, 7, 60, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12,
	76, 10, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 2,
	4, 3, 2, 8, 9, 3, 2, 15, 16, 2, 77, 2, 27, 3, 2, 2, 2, 4, 34, 3, 2, 2,
	2, 6, 36, 3, 2, 2, 2, 8, 47, 3, 2, 2, 2, 10, 52, 3, 2, 2, 2, 12, 54, 3,
	2, 2, 2, 14, 63, 3, 2, 2, 2, 16, 65, 3, 2, 2, 2, 18, 67, 3, 2, 2, 2, 20,
	69, 3, 2, 2, 2, 22, 75, 3, 2, 2, 2, 24, 26, 5, 4, 3, 2, 25, 24, 3, 2, 2,
	2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3, 2, 2, 2, 28, 3, 3,
	2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 35, 5, 6, 4, 2, 31, 35, 5, 22, 12, 2,
	32, 35, 5, 8, 5, 2, 33, 35, 5, 12, 7, 2, 34, 30, 3, 2, 2, 2, 34, 31, 3,
	2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35, 5, 3, 2, 2, 2, 36,
	45, 9, 2, 2, 2, 37, 41, 7, 3, 2, 2, 38, 40, 5, 22, 12, 2, 39, 38, 3, 2,
	2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 44,
	3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 46, 7, 4, 2, 2, 45, 37, 3, 2, 2, 2,
	45, 46, 3, 2, 2, 2, 46, 7, 3, 2, 2, 2, 47, 48, 7, 5, 2, 2, 48, 49, 7, 9,
	2, 2, 49, 9, 3, 2, 2, 2, 50, 53, 5, 6, 4, 2, 51, 53, 5, 22, 12, 2, 52,
	50, 3, 2, 2, 2, 52, 51, 3, 2, 2, 2, 53, 11, 3, 2, 2, 2, 54, 58, 7, 6, 2,
	2, 55, 57, 5, 10, 6, 2, 56, 55, 3, 2, 2, 2, 57, 60, 3, 2, 2, 2, 58, 56,
	3, 2, 2, 2, 58, 59, 3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2,
	61, 62, 7, 7, 2, 2, 62, 13, 3, 2, 2, 2, 63, 64, 7, 10, 2, 2, 64, 15, 3,
	2, 2, 2, 65, 66, 7, 12, 2, 2, 66, 17, 3, 2, 2, 2, 67, 68, 7, 13, 2, 2,
	68, 19, 3, 2, 2, 2, 69, 70, 9, 3, 2, 2, 70, 21, 3, 2, 2, 2, 71, 76, 5,
	14, 8, 2, 72, 76, 5, 16, 9, 2, 73, 76, 5, 18, 10, 2, 74, 76, 5, 20, 11,
	2, 75, 71, 3, 2, 2, 2, 75, 72, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 74,
	3, 2, 2, 2, 76, 23, 3, 2, 2, 2, 9, 27, 34, 41, 45, 52, 58, 75,
}
var literalNames = []string{
	"", "'['", "']'", "'$'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "OPS", "NAME", "INTEGER", "DECIMAL_INTEGER", "FLOAT_NUMBER",
	"STRING", "OP", "TRUE", "FALSE", "BLOCK_COMMENT", "WhiteSpace", "NewLine",
}

var ruleNames = []string{
	"expressions", "root_term", "fun", "vars", "dblock_term", "dblock", "integer_term",
	"float_term", "string_term", "boolean_term", "term",
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
	ThreadComputationParserOPS             = 6
	ThreadComputationParserNAME            = 7
	ThreadComputationParserINTEGER         = 8
	ThreadComputationParserDECIMAL_INTEGER = 9
	ThreadComputationParserFLOAT_NUMBER    = 10
	ThreadComputationParserSTRING          = 11
	ThreadComputationParserOP              = 12
	ThreadComputationParserTRUE            = 13
	ThreadComputationParserFALSE           = 14
	ThreadComputationParserBLOCK_COMMENT   = 15
	ThreadComputationParserWhiteSpace      = 16
	ThreadComputationParserNewLine         = 17
)

// ThreadComputationParser rules.
const (
	ThreadComputationParserRULE_expressions  = 0
	ThreadComputationParserRULE_root_term    = 1
	ThreadComputationParserRULE_fun          = 2
	ThreadComputationParserRULE_vars         = 3
	ThreadComputationParserRULE_dblock_term  = 4
	ThreadComputationParserRULE_dblock       = 5
	ThreadComputationParserRULE_integer_term = 6
	ThreadComputationParserRULE_float_term   = 7
	ThreadComputationParserRULE_string_term  = 8
	ThreadComputationParserRULE_boolean_term = 9
	ThreadComputationParserRULE_term         = 10
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

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__2)|(1<<ThreadComputationParserT__3)|(1<<ThreadComputationParserOPS)|(1<<ThreadComputationParserNAME)|(1<<ThreadComputationParserINTEGER)|(1<<ThreadComputationParserFLOAT_NUMBER)|(1<<ThreadComputationParserSTRING)|(1<<ThreadComputationParserTRUE)|(1<<ThreadComputationParserFALSE))) != 0 {
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
	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserOPS, ThreadComputationParserNAME:
		{
			p.SetState(28)
			p.Fun()
		}

	case ThreadComputationParserINTEGER, ThreadComputationParserFLOAT_NUMBER, ThreadComputationParserSTRING, ThreadComputationParserTRUE, ThreadComputationParserFALSE:
		{
			p.SetState(29)
			p.Term()
		}

	case ThreadComputationParserT__2:
		{
			p.SetState(30)
			p.Vars()
		}

	case ThreadComputationParserT__3:
		{
			p.SetState(31)
			p.Dblock()
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

	// Get_term returns the _term rule contexts.
	Get_term() ITermContext

	// Set_term sets the _term rule contexts.
	Set_term(ITermContext)

	// GetParam returns the param rule context list.
	GetParam() []ITermContext

	// SetParam sets the param rule context list.
	SetParam([]ITermContext)

	// IsFunContext differentiates from other interfaces.
	IsFunContext()
}

type FunContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	fname  antlr.Token
	_term  ITermContext
	param  []ITermContext
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

func (s *FunContext) Get_term() ITermContext { return s._term }

func (s *FunContext) Set_term(v ITermContext) { s._term = v }

func (s *FunContext) GetParam() []ITermContext { return s.param }

func (s *FunContext) SetParam(v []ITermContext) { s.param = v }

func (s *FunContext) NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserNAME, 0)
}

func (s *FunContext) OPS() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserOPS, 0)
}

func (s *FunContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *FunContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
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
		p.SetState(34)

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
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThreadComputationParserT__0 {
		{
			p.SetState(35)
			p.Match(ThreadComputationParserT__0)
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserINTEGER)|(1<<ThreadComputationParserFLOAT_NUMBER)|(1<<ThreadComputationParserSTRING)|(1<<ThreadComputationParserTRUE)|(1<<ThreadComputationParserFALSE))) != 0 {
			{
				p.SetState(36)

				var _x = p.Term()

				localctx.(*FunContext)._term = _x
			}
			localctx.(*FunContext).param = append(localctx.(*FunContext).param, localctx.(*FunContext)._term)

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(42)
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
	p.EnterRule(localctx, 6, ThreadComputationParserRULE_vars)

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
		p.SetState(45)
		p.Match(ThreadComputationParserT__2)
	}
	{
		p.SetState(46)

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
	p.EnterRule(localctx, 8, ThreadComputationParserRULE_dblock_term)

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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserOPS, ThreadComputationParserNAME:
		{
			p.SetState(48)
			p.Fun()
		}

	case ThreadComputationParserINTEGER, ThreadComputationParserFLOAT_NUMBER, ThreadComputationParserSTRING, ThreadComputationParserTRUE, ThreadComputationParserFALSE:
		{
			p.SetState(49)
			p.Term()
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
	p.EnterRule(localctx, 10, ThreadComputationParserRULE_dblock)
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
		p.SetState(52)
		p.Match(ThreadComputationParserT__3)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserOPS)|(1<<ThreadComputationParserNAME)|(1<<ThreadComputationParserINTEGER)|(1<<ThreadComputationParserFLOAT_NUMBER)|(1<<ThreadComputationParserSTRING)|(1<<ThreadComputationParserTRUE)|(1<<ThreadComputationParserFALSE))) != 0 {
		{
			p.SetState(53)

			var _x = p.Dblock_term()

			localctx.(*DblockContext)._dblock_term = _x
		}
		localctx.(*DblockContext).param = append(localctx.(*DblockContext).param, localctx.(*DblockContext)._dblock_term)

		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(59)
		p.Match(ThreadComputationParserT__4)
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
	p.EnterRule(localctx, 12, ThreadComputationParserRULE_integer_term)

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
		p.SetState(61)

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
	p.EnterRule(localctx, 14, ThreadComputationParserRULE_float_term)

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
		p.SetState(63)

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
	p.EnterRule(localctx, 16, ThreadComputationParserRULE_string_term)

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
		p.SetState(65)

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
	p.EnterRule(localctx, 18, ThreadComputationParserRULE_boolean_term)
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
		p.SetState(67)

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
	p.EnterRule(localctx, 20, ThreadComputationParserRULE_term)

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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserINTEGER:
		{
			p.SetState(69)
			p.Integer_term()
		}

	case ThreadComputationParserFLOAT_NUMBER:
		{
			p.SetState(70)
			p.Float_term()
		}

	case ThreadComputationParserSTRING:
		{
			p.SetState(71)
			p.String_term()
		}

	case ThreadComputationParserTRUE, ThreadComputationParserFALSE:
		{
			p.SetState(72)
			p.Boolean_term()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
