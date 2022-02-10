package ThreadComputation

import (
  "os"
  log "github.com/sirupsen/logrus"
  "github.com/antlr/antlr4/runtime/Go/antlr"
  "github.com/gammazero/deque"
  "github.com/vulogov/ThreadComputation/parser"
)

type TCExecListener struct {
  *parser.BaseThreadComputationListener
  TC          *TCstate
}

type TCstate struct {
  TC          *TCstate
  errors       int
  InAttr       bool
  Attrs        deque.Deque
  Res          deque.Deque
}

type tcExecErrorListener struct {
	antlr.ErrorListener
  TC     *TCstate
	code   *string
	errors int
}

func Init() *TCstate {
  log.SetOutput(os.Stderr)
  log.SetLevel(log.InfoLevel)
  tc := &TCstate{
    InAttr:  false,
    errors:  0,
  }
  return tc
}

func (tc *TCstate) Eval(code string) *TCstate {
  errorListener := new(tcExecErrorListener)
  errorListener.code = &code
  _input := antlr.NewInputStream(code)
  lexer := parser.NewThreadComputationLexer(_input)
  stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
  p := parser.NewThreadComputationParser(stream)
  p.RemoveErrorListeners()
	p.AddErrorListener(errorListener)
  listener := new(TCExecListener)
  listener.TC = tc
  errorListener.TC = tc
  tc.errors = errorListener.errors
  if errorListener.errors > 0 {
		log.Fatalf("Lexer errors detected: %v", errorListener.errors)
		return nil
	}
  antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
  tc.errors = errorListener.errors
  if errorListener.errors > 0 {
		log.Fatalf("Errors detected: %v", errorListener.errors)
		return nil
	}
  return tc
}

func (tc *TCstate) Errors() int {
  return tc.errors
}

func (tc *TCstate) Get() interface{} {
  if tc.Res.Len() == 0 {
    return nil
  }
  return tc.Res.PopFront()
}

func (l *tcExecErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	log.Errorf("Syntax error line=%v, column=%v : %v", line, column, msg)
	l.errors += 1
}
func (l *tcExecErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	log.Errorf("Ambiguity Error")
	l.errors += 1
}
func (l *tcExecErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	log.Errorf("Attempting in Full Context")
	l.errors += 1
}
func (l *tcExecErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	log.Errorf("Context sensitivity error")
	l.errors += 1
}
