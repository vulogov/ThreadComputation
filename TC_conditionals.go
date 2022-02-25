package ThreadComputation

import (
  log "github.com/sirupsen/logrus"
  "github.com/vulogov/ThreadComputation/parser"
)

func TCRunConditional(l *TCExecListener, conditional bool) {
  var func_name string
  if l.TC.UFStack.Len() > 0 {
    func_name = l.TC.UFNStack.Front().(string)
    code := l.TC.UFStack.PopFront()
    l.TC.UFNStack.PopFront()
    if l.TC.Ready() {
      e := l.TC.Get()
      switch e.(type) {
      case bool:
        if e.(bool) == conditional {
          log.Debugf("pre-condition met for conditional: %v=%v %v", e, conditional, func_name)
          return
        }
      }
      l.TC.Res.Set(e)
      for {
        log.Debugf("Running conditional: %v", func_name)
        l.TC.Eval(code.(string))
        e = l.TC.Get()
        log.Debugf("Get value from stack for conditional: %T", e)
        switch e.(type) {
        case bool:
          if e.(bool) == conditional {
            log.Debugf("Condition met for conditional: %v=%v %v", e, conditional, func_name)
            return
          }
        default:
          log.Debugf("Return from conditional due to non-boolean data in stack: %T", e)
          ReturnFromFunction(l, "conditional", e)
          return
        }
      }
    }
  }
  return
}

func (l *TCExecListener) EnterTrueblock(c *parser.TrueblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  log.Debug("Entering TRUE block")
  MakeTempFun(l)
}

func (l *TCExecListener) EnterFalseblock(c *parser.FalseblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  log.Debug("Entering FALSE block")
  MakeTempFun(l)
}

func (l *TCExecListener) ExitTrueblock(c *parser.TrueblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  TCRunConditional(l, true)
  log.Debug("Exiting TRUE block")
}

func (l *TCExecListener) ExitFalseblock(c *parser.FalseblockContext) {
  if l.TC.Errors() > 0 {
    return
  }
  TCRunConditional(l, false)
  log.Debug("Exiting TRUE block")
}