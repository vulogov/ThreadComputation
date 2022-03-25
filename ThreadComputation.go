package ThreadComputation

import (
  "io"
  "os"
  "fmt"
  "sync"
  "github.com/lrita/cmap"
  log "github.com/sirupsen/logrus"
  "github.com/antlr/antlr4/runtime/Go/antlr"
  "github.com/gammazero/deque"
  "github.com/deckarep/golang-set"
  "github.com/google/uuid"
  "github.com/vulogov/ThreadComputation/parser"
  "github.com/loveleshsharma/gohive"
)

var Vars      cmap.Cmap
var Functions cmap.Cmap
var Commands  cmap.Cmap
var Callbacks cmap.Cmap
var VERSION = "1.15"

type TCExecListener struct {
  *parser.BaseThreadComputationListener
  TC          *TCstate
}

type TCstate struct {
  TC          *TCstate
  errors       int
  errmsg       string
  HandleErr    bool
  InAttr       int          // How deep we are in attributes
  InRef        int          // If we are in reference
  Attrs       *TwoStack     // Creating attributes
  EvAttrs      deque.Deque  // Evaluating attributes
  Res         *TwoStack     // Main stack
  ResNames     deque.Deque  // stack of stack names
  ResN         mapset.Set   // set of stack names
  Vars         cmap.Cmap    // variables
  Functions    cmap.Cmap    // Functions
  UFunctions   cmap.Cmap    // User Functions
  Commands     cmap.Cmap    // Commands
  UFStack      deque.Deque  // Stack of User-defined functions
  UFNStack     deque.Deque  // Stack of names User-defined functions
  UFNB         int          // balancing []
  CStack       deque.Deque  // Conditionals stack
  StackList    cmap.Cmap    // Reference to stacks
  StackChan    cmap.Cmap    // Reference to stack channels
  Wg           sync.WaitGroup // Global wait group
  Pool        *gohive.PoolService // Global execution pool
}

type tcExecErrorListener struct {
	antlr.ErrorListener
  TC     *TCstate
	code   *string
	errors int
}

func Init() *TCstate {
  var pool_size interface{}
  out, err := GetVariable("tc.Logoutput")
  if err != nil {
    log.SetOutput(os.Stderr)
  } else {
    log.SetOutput(out.(io.Writer))
  }
  lvl, err :=  GetVariable("tc.Debuglevel")
  if err != nil {
    lvl = "info"
  }
  switch lvl {
  case "trace":
    log.SetLevel(log.TraceLevel)
  case "debug":
    log.SetLevel(log.DebugLevel)
  case "info":
    log.SetLevel(log.InfoLevel)
  case "warning":
    log.SetLevel(log.WarnLevel)
  case "error":
    log.SetLevel(log.ErrorLevel)
  case "fatal":
    log.SetLevel(log.FatalLevel)
  default:
    log.SetLevel(log.InfoLevel)
  }
  pool_size, err = GetVariable("tc.Poolsize")
  if err != nil {
    pool_size = 25
  }
  log.Debug("Creating TC")
  tc := &TCstate{
    InAttr:  0,
    InRef:   0,
    errors:  0,
    HandleErr: false,
    UFNB:    0,
    Res:     InitTS(),
    Attrs:   InitTS(),
    ResN:    mapset.NewSet(),
    Pool:    gohive.NewFixedSizePool(pool_size.(int)),
  }
  tc.AddNewStack(uuid.NewString())
  return tc
}

func (tc *TCstate) ClearErrors() {
  log.Debug("Clearing errors")
  tc.errors = 0
  tc.errmsg = ""
}

func (tc *TCstate) Eval(code string) *TCstate {
  log.Debugf("Eval: %v", code)
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

func (tc *TCstate) GoEval(code string) *TCstate {
  tc.Wg.Add(1)
  task := func() {
    defer tc.Wg.Done()
    tc.Eval(code)
  }
  tc.Pool.Submit(task)
  return tc
}

func (tc *TCstate) Errors() int {
  if tc.HandleErr {
    return 0
  }
  return tc.errors
}

func (tc *TCstate) TrueErrors() int {
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

func (tc *TCstate) Set(x interface{}) {
  if tc.Res.GLen() == 0 {
    return
  }
  tc.Res.Set(x)
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

// func (tc *TCstate) CurrentFunctionName() string {
//   if tc.FNStack.Len() > 0 {
//     return tc.FNStack.Front().(string)
//   } else {
//     return "#MAIN"
//   }
// }



func (tc *TCstate) SetError(msg string, args ...interface{}) {
  tc.errmsg = fmt.Sprintf(msg, args...)
  tc.errors += 1
  log.Errorf(tc.errmsg)
}

func (l *TCExecListener) SetError(msg string, args ...interface{}) {
  l.TC.SetError(msg, args...)
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
  initStdVars()
}
