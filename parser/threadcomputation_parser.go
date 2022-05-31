// Code generated from ThreadComputation.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ThreadComputation

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type ThreadComputationParser struct {
	*antlr.BaseParser
}

var threadcomputationParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func threadcomputationParserInit() {
	staticData := &threadcomputationParserStaticData
	staticData.literalNames = []string{
		"", "'['", "':'", "';;'", "']'", "'{'", "'}'", "'#'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "FUNC_NAME", "OPS", "NAME", "INTEGER",
		"DECIMAL_INTEGER", "FLOAT_NUMBER", "STRING", "OP", "MOD", "BLOCK_COMMENT",
		"WhiteSpace", "NewLine",
	}
	staticData.ruleNames = []string{
		"expressions", "root_term", "ns", "fun", "val", "ns_term", "fun_term",
		"pos_term",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 75, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 5, 0, 18, 8, 0, 10, 0, 12,
		0, 21, 9, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 27, 8, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 5, 2, 33, 8, 2, 10, 2, 12, 2, 36, 9, 2, 1, 2, 1, 2, 1, 3, 3, 3, 41,
		8, 3, 1, 3, 1, 3, 1, 3, 5, 3, 46, 8, 3, 10, 3, 12, 3, 49, 9, 3, 1, 3, 3,
		3, 52, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 57, 8, 4, 1, 4, 1, 4, 1, 5, 1, 5,
		1, 5, 1, 5, 3, 5, 65, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6, 70, 8, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 0, 0, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 0, 80, 0, 19, 1,
		0, 0, 0, 2, 26, 1, 0, 0, 0, 4, 28, 1, 0, 0, 0, 6, 40, 1, 0, 0, 0, 8, 53,
		1, 0, 0, 0, 10, 64, 1, 0, 0, 0, 12, 69, 1, 0, 0, 0, 14, 71, 1, 0, 0, 0,
		16, 18, 3, 2, 1, 0, 17, 16, 1, 0, 0, 0, 18, 21, 1, 0, 0, 0, 19, 17, 1,
		0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 1, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 22,
		27, 3, 6, 3, 0, 23, 27, 3, 4, 2, 0, 24, 27, 3, 14, 7, 0, 25, 27, 3, 8,
		4, 0, 26, 22, 1, 0, 0, 0, 26, 23, 1, 0, 0, 0, 26, 24, 1, 0, 0, 0, 26, 25,
		1, 0, 0, 0, 27, 3, 1, 0, 0, 0, 28, 29, 5, 1, 0, 0, 29, 30, 5, 8, 0, 0,
		30, 34, 5, 2, 0, 0, 31, 33, 3, 10, 5, 0, 32, 31, 1, 0, 0, 0, 33, 36, 1,
		0, 0, 0, 34, 32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 37, 1, 0, 0, 0, 36,
		34, 1, 0, 0, 0, 37, 38, 5, 3, 0, 0, 38, 5, 1, 0, 0, 0, 39, 41, 5, 16, 0,
		0, 40, 39, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0, 42, 51,
		5, 8, 0, 0, 43, 47, 5, 1, 0, 0, 44, 46, 3, 12, 6, 0, 45, 44, 1, 0, 0, 0,
		46, 49, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0, 48, 50, 1,
		0, 0, 0, 49, 47, 1, 0, 0, 0, 50, 52, 5, 4, 0, 0, 51, 43, 1, 0, 0, 0, 51,
		52, 1, 0, 0, 0, 52, 7, 1, 0, 0, 0, 53, 54, 5, 5, 0, 0, 54, 56, 3, 6, 3,
		0, 55, 57, 5, 8, 0, 0, 56, 55, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58,
		1, 0, 0, 0, 58, 59, 5, 6, 0, 0, 59, 9, 1, 0, 0, 0, 60, 65, 3, 6, 3, 0,
		61, 65, 3, 4, 2, 0, 62, 65, 3, 14, 7, 0, 63, 65, 3, 8, 4, 0, 64, 60, 1,
		0, 0, 0, 64, 61, 1, 0, 0, 0, 64, 62, 1, 0, 0, 0, 64, 63, 1, 0, 0, 0, 65,
		11, 1, 0, 0, 0, 66, 70, 3, 6, 3, 0, 67, 70, 3, 14, 7, 0, 68, 70, 3, 8,
		4, 0, 69, 66, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 13,
		1, 0, 0, 0, 71, 72, 5, 7, 0, 0, 72, 73, 5, 8, 0, 0, 73, 15, 1, 0, 0, 0,
		9, 19, 26, 34, 40, 47, 51, 56, 64, 69,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// ThreadComputationParserInit initializes any static state used to implement ThreadComputationParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewThreadComputationParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ThreadComputationParserInit() {
	staticData := &threadcomputationParserStaticData
	staticData.once.Do(threadcomputationParserInit)
}

// NewThreadComputationParser produces a new parser instance for the optional input antlr.TokenStream.
func NewThreadComputationParser(input antlr.TokenStream) *ThreadComputationParser {
	ThreadComputationParserInit()
	this := new(ThreadComputationParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &threadcomputationParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
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
	ThreadComputationParserFUNC_NAME       = 8
	ThreadComputationParserOPS             = 9
	ThreadComputationParserNAME            = 10
	ThreadComputationParserINTEGER         = 11
	ThreadComputationParserDECIMAL_INTEGER = 12
	ThreadComputationParserFLOAT_NUMBER    = 13
	ThreadComputationParserSTRING          = 14
	ThreadComputationParserOP              = 15
	ThreadComputationParserMOD             = 16
	ThreadComputationParserBLOCK_COMMENT   = 17
	ThreadComputationParserWhiteSpace      = 18
	ThreadComputationParserNewLine         = 19
)

// ThreadComputationParser rules.
const (
	ThreadComputationParserRULE_expressions = 0
	ThreadComputationParserRULE_root_term   = 1
	ThreadComputationParserRULE_ns          = 2
	ThreadComputationParserRULE_fun         = 3
	ThreadComputationParserRULE_val         = 4
	ThreadComputationParserRULE_ns_term     = 5
	ThreadComputationParserRULE_fun_term    = 6
	ThreadComputationParserRULE_pos_term    = 7
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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRoot_termContext); ok {
			len++
		}
	}

	tst := make([]IRoot_termContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRoot_termContext); ok {
			tst[i] = t.(IRoot_termContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionsContext) Root_term(i int) IRoot_termContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoot_termContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__0)|(1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserT__6)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
		{
			p.SetState(16)
			p.Root_term()
		}

		p.SetState(21)
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *Root_termContext) Ns() INsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *Root_termContext) Pos_term() IPos_termContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPos_termContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPos_termContext)
}

func (s *Root_termContext) Val() IValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValContext)
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
	p.SetState(26)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserFUNC_NAME, ThreadComputationParserMOD:
		{
			p.SetState(22)
			p.Fun()
		}

	case ThreadComputationParserT__0:
		{
			p.SetState(23)
			p.Ns()
		}

	case ThreadComputationParserT__6:
		{
			p.SetState(24)
			p.Pos_term()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(25)
			p.Val()
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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INs_termContext); ok {
			len++
		}
	}

	tst := make([]INs_termContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INs_termContext); ok {
			tst[i] = t.(INs_termContext)
			i++
		}
	}

	return tst
}

func (s *NsContext) Ns_term(i int) INs_termContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INs_termContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
		p.SetState(28)
		p.Match(ThreadComputationParserT__0)
	}
	{
		p.SetState(29)

		var _m = p.Match(ThreadComputationParserFUNC_NAME)

		localctx.(*NsContext).nsname = _m
	}
	{
		p.SetState(30)
		p.Match(ThreadComputationParserT__1)
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__0)|(1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserT__6)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
		{
			p.SetState(31)

			var _x = p.Ns_term()

			localctx.(*NsContext)._ns_term = _x
		}
		localctx.(*NsContext).param = append(localctx.(*NsContext).param, localctx.(*NsContext)._ns_term)

		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(37)
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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFun_termContext); ok {
			len++
		}
	}

	tst := make([]IFun_termContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFun_termContext); ok {
			tst[i] = t.(IFun_termContext)
			i++
		}
	}

	return tst
}

func (s *FunContext) Fun_term(i int) IFun_termContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFun_termContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThreadComputationParserMOD {
		{
			p.SetState(39)

			var _m = p.Match(ThreadComputationParserMOD)

			localctx.(*FunContext).mod = _m
		}

	}
	{
		p.SetState(42)

		var _m = p.Match(ThreadComputationParserFUNC_NAME)

		localctx.(*FunContext).fname = _m
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(43)
			p.Match(ThreadComputationParserT__0)
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThreadComputationParserT__4)|(1<<ThreadComputationParserT__6)|(1<<ThreadComputationParserFUNC_NAME)|(1<<ThreadComputationParserMOD))) != 0 {
			{
				p.SetState(44)

				var _x = p.Fun_term()

				localctx.(*FunContext)._fun_term = _x
			}
			localctx.(*FunContext).param = append(localctx.(*FunContext).param, localctx.(*FunContext)._fun_term)

			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(50)
			p.Match(ThreadComputationParserT__3)
		}

	}

	return localctx
}

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetP returns the p token.
	GetP() antlr.Token

	// SetP sets the p token.
	SetP(antlr.Token)

	// GetFname returns the fname rule contexts.
	GetFname() IFunContext

	// SetFname sets the fname rule contexts.
	SetFname(IFunContext)

	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	fname  IFunContext
	p      antlr.Token
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThreadComputationParserRULE_val
	return p
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThreadComputationParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) GetP() antlr.Token { return s.p }

func (s *ValContext) SetP(v antlr.Token) { s.p = v }

func (s *ValContext) GetFname() IFunContext { return s.fname }

func (s *ValContext) SetFname(v IFunContext) { s.fname = v }

func (s *ValContext) Fun() IFunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *ValContext) FUNC_NAME() antlr.TerminalNode {
	return s.GetToken(ThreadComputationParserFUNC_NAME, 0)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.EnterVal(s)
	}
}

func (s *ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThreadComputationListener); ok {
		listenerT.ExitVal(s)
	}
}

func (p *ThreadComputationParser) Val() (localctx IValContext) {
	this := p
	_ = this

	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ThreadComputationParserRULE_val)
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
		p.SetState(53)
		p.Match(ThreadComputationParserT__4)
	}
	{
		p.SetState(54)

		var _x = p.Fun()

		localctx.(*ValContext).fname = _x
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThreadComputationParserFUNC_NAME {
		{
			p.SetState(55)

			var _m = p.Match(ThreadComputationParserFUNC_NAME)

			localctx.(*ValContext).p = _m
		}

	}
	{
		p.SetState(58)
		p.Match(ThreadComputationParserT__5)
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *Ns_termContext) Ns() INsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INsContext)
}

func (s *Ns_termContext) Pos_term() IPos_termContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPos_termContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPos_termContext)
}

func (s *Ns_termContext) Val() IValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValContext)
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
	p.EnterRule(localctx, 10, ThreadComputationParserRULE_ns_term)

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
	p.SetState(64)
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

	case ThreadComputationParserT__6:
		{
			p.SetState(62)
			p.Pos_term()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(63)
			p.Val()
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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunContext)
}

func (s *Fun_termContext) Pos_term() IPos_termContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPos_termContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPos_termContext)
}

func (s *Fun_termContext) Val() IValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValContext)
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
	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThreadComputationParserFUNC_NAME, ThreadComputationParserMOD:
		{
			p.SetState(66)
			p.Fun()
		}

	case ThreadComputationParserT__6:
		{
			p.SetState(67)
			p.Pos_term()
		}

	case ThreadComputationParserT__4:
		{
			p.SetState(68)
			p.Val()
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
		p.SetState(71)
		p.Match(ThreadComputationParserT__6)
	}
	{
		p.SetState(72)

		var _m = p.Match(ThreadComputationParserFUNC_NAME)

		localctx.(*Pos_termContext).pname = _m
	}

	return localctx
}
