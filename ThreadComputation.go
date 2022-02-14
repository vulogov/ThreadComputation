package ThreadComputation

import (
  "os"
  "fmt"
  "github.com/lrita/cmap"
  log "github.com/sirupsen/logrus"
  "github.com/antlr/antlr4/runtime/Go/antlr"
  "github.com/gammazero/deque"
  "github.com/vulogov/ThreadComputation/parser"
)

var Vars      cmap.Cmap
var Functions cmap.Cmap

type TCExecListener struct {
  *parser.BaseThreadComputationListener
  TC          *TCstate
}

type TCstate struct {
  TC          *TCstate
  errors       int
  errmsg       string
  InAttr       int
  InRef        int
  Attrs       *TwoStack
  Res         *TwoStack
  FNStack      deque.Deque
  Vars         cmap.Cmap
  Functions    cmap.Cmap
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
    InAttr:  0,
    InRef:   0,
    errors:  0,
    Res:     InitTS(),
    Attrs:   InitTS(),
  }
  tc.Res.Add()
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
  if errorListener.errors > 0 {
    tc.errors = errorListener.errors
  }
  if errorListener.errors > 0 {
		log.Fatalf("Lexer errors detected: %v", errorListener.errors)
		return nil
	}
  antlr.ParseTreeWalkerDefault.Walk(listener, p.Expressions())
  if errorListener.errors > 0 {
    tc.errors = errorListener.errors
  }
  if errorListener.errors > 0 {
		log.Fatalf("Errors detected: %v", errorListener.errors)
		return nil
	}
  return tc
}

func (tc *TCstate) Errors() int {
  return tc.errors
}

func (tc *TCstate) Error() string {
  return tc.errmsg
}

func (tc *TCstate) Get() interface{} {
  if tc.Res.Len() == 0 {
    return nil
  }
  res, err := tc.Res.Take()
  if err == nil {
    return res
  }
  return nil
}

func (tc *TCstate) GetAsString() string {
  res := tc.Get()
  if res != nil {
    return fmt.Sprintf("%v", res)
  }
  return ""
}

func (tc *TCstate) Ready() bool {
  if tc.Res.Len() == 0 {
    return false
  }
  return true
}

func (tc *TCstate) HaveAttrs() bool {
  if tc.Attrs.Len() == 0 {
    return false
  }
  return true
}

func (tc *TCstate) CurrentFunctionName() string {
  if tc.FNStack.Len() > 0 {
    return tc.FNStack.Front().(string)
  } else {
    return "#MAIN"
  }
}

func (tc *TCstate) SetFunction(name string, fun TCFun) {
  tc.Functions.Delete(name)
  tc.Functions.Store(name, fun)
}

func (tc *TCstate) SetVariable(name string, data interface{}) {
  tc.Vars.Delete(name)
  tc.Vars.Store(name, data)
}

func (l *tcExecErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
  msgout := fmt.Sprintf("Syntax error line=%v, column=%v : %v", line, column, msg)
  log.Errorf(msgout)
  l.TC.errmsg = msgout
	l.errors += 1
}
func (l *tcExecErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
  msgout := fmt.Sprintf("Ambiguity Error")
  log.Errorf(msgout)
  l.TC.errmsg = msgout
	l.errors += 1
}
func (l *tcExecErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
  msgout := fmt.Sprintf("Attempting in Full Context")
  log.Errorf(msgout)
  l.TC.errmsg = msgout
	l.errors += 1
}
func (l *tcExecErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
  msgout := fmt.Sprintf("Context sensitivity error")
  log.Errorf(msgout)
  l.TC.errmsg = msgout
	l.errors += 1
}

func init() {
  initStdlib()
}
